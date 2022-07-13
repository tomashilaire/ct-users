package main

import (
	"flag"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"log"
	"net"
	"os"
	"root/internal/core/services/filesrv"
	"root/internal/core/services/userssrv"
	"root/internal/handlers/filesprotohdl"
	"root/internal/handlers/usersprotohdl"
	"root/internal/interceptors/authgrpcintrcp"
	"root/internal/repositories/filess3repo"
	"root/internal/repositories/usersmongorepo"
	"root/pb"
	"root/pkg/security"
	"root/pkg/uidgen"
	"root/pkg/validators"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port  int
	local bool
)

func init() {
	flag.IntVar(&port, "port", 9092, "gRCP service port")
	flag.BoolVar(&local, "local", true, "run service local")
	flag.Parse()
}

func accessibleMethods() map[string][]string {
	const authServicePath = "/pb.Authentication/"

	return map[string][]string{
		authServicePath + "Authenticate": {},
	}
}

func tracerProvider(url string) (*trace.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exp),
		// Record information about this application in a Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(os.Getenv("SERVICE_NAME")),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "demo"),
		)),
	)
	return tp, nil
}

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Println("Unable to retrieve env variables", err)
		}
	}

	// instance repository, service and handlers -> register handlers
	fr := filess3repo.NewFilesRepository()
	fs := filesrv.NewService(fr, uidgen.New())
	fh := filesprotohdl.NewProtoHandler(fs)

	// instance repository, service and handlers -> register handlers
	ur := usersmongorepo.NewUsersRepository()
	defer ur.Disconnect()
	us := userssrv.NewService(ur, uidgen.New(), validators.NewValidators(), security.NewSecurity())
	uh := usersprotohdl.NewProtoHandler(us)

	// run server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Println("Unable to listen", err)
		os.Exit(1)
	}
	interceptor := authgrpcintrcp.NewAuthInterceptor(security.NewSecurity(), accessibleMethods())

	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}

	// register server
	gs := grpc.NewServer(serverOptions...)
	reflection.Register(gs)

	pb.RegisterFilesServer(gs, fh)
	pb.RegisterAuthenticationServer(gs, uh)

	log.Println(fmt.Sprintf("grpc service running on [::]:%d", port))

	gs.Serve(listener)
}
