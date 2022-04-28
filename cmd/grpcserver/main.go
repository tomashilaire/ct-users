package main

import (
	"flag"
	"fmt"
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
