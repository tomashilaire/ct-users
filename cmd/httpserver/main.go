package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"test/cmd/httpserver/routes"
	"test/internal/handlers/testhttphdl"
	"test/pb"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var (
	grcpPort int
	grcpAddr string
)

func init() {
	flag.IntVar(&grcpPort, "port", 9090, "HTTP service port")
	flag.StringVar(&grcpAddr, "test_addr", "localhost:9092", "test service address")
	flag.Parse()
}

func main() {

	conn, err := grpc.Dial(grcpAddr, grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	// to check
	tc := pb.NewTestClient(conn)
	th := testhttphdl.NewHTTPHandler(tc)
	tr := routes.NewTestRoutes(th)

	router := mux.NewRouter().StrictSlash(true)
	routes.Install(router, tr)

	log.Printf("API service running on [::]:%d\n", grcpPort)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", grcpPort), routes.WithCORS(router)))
}
