package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"test/internal/core/services/testsrv"
	"test/internal/handlers/testprotohdl"
	"test/internal/repositories/testmongorepo"
	"test/pb"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port  int
	local bool
)

func init() {
	flag.IntVar(&port, "port", 9092, "test service port")
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
	cfg := testmongorepo.NewConfig()
	db, err := testmongorepo.NewConnection(cfg)
	if err != nil {
		log.Println("Unable to connect", err)
	}
	defer db.Disconnect()

	// instance repository, service and handlers -> register handlers
	tr := testmongorepo.NewTestRepository(db)
	ts := testsrv.NewService(tr)
	th := testprotohdl.NewProtoHandler(ts)

	// run server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Println("Unable to listen", err)
		os.Exit(1)
	}

	// register server
	gs := grpc.NewServer()
	reflection.Register(gs)

	pb.RegisterTestServer(gs, th)

	log.Println(fmt.Sprintf("Service running on [::]:%d", port))

	gs.Serve(listener)
}