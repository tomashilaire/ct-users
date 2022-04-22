package main

import (
	"entity/internal/core/services/entitysrv"
	"entity/internal/handlers/entityprotohdl"
	"entity/internal/repositories/entitymongorepo"
	"entity/pb"
	"entity/pkg/uidgen"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

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

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Println("Unable to retrieve env variables", err)
		}
	}

	// db config and conn
	cfg := entitymongorepo.NewConfig()
	db, err := entitymongorepo.NewConnection(cfg)
	if err != nil {
		log.Println("Unable to connect", err)
	}
	defer db.Disconnect()

	// instance repository, service and handlers -> register handlers
	tr := entitymongorepo.NewEntityRepository(db)
	ts := entitysrv.NewService(tr, uidgen.New())
	th := entityprotohdl.NewProtoHandler(ts)

	// run server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Println("Unable to listen", err)
		os.Exit(1)
	}

	// register server
	gs := grpc.NewServer()
	reflection.Register(gs)

	pb.RegisterEntityServer(gs, th)

	log.Println(fmt.Sprintf("grpc service running on [::]:%d", port))

	gs.Serve(listener)
}
