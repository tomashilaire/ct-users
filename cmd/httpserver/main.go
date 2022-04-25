package main

import (
	"context"
	"entity/internal/core/services/entitysrv"
	"entity/internal/handlers/entityhttphdl"
	"entity/internal/repositories/entitymongorepo"
	"entity/pkg/uidgen"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	local bool
)

func init() {
	flag.BoolVar(&local, "local", true, "run service local")
	flag.Parse()
}

func main() {
	// retrieve env variables
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

	// instance repository and service -> register handlers
	tr := entitymongorepo.NewEntityRepository(db)
	ts := entitysrv.NewService(tr, uidgen.New())
	th := entityhttphdl.NewHTTPHandler(ts)

	r := mux.NewRouter()

	// establish subrouters
	getRouter := r.Methods("GET").Subrouter()
	getRouter.HandleFunc("/entities", th.GetAllEntities)
	getRouter.HandleFunc("/entities/{id:[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}}", th.GetEntity)

	postRouter := r.Methods("POST").Subrouter()
	postRouter.HandleFunc("/entities", th.PostEntity)

	putRouter := r.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/entities/{id:[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}}", th.PutEntity)

	deleteRouter := r.Methods("DELETE").Subrouter()
	deleteRouter.HandleFunc("/entities/{id:[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}}", th.DeleteEntity)

	// serve docs endpoint
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// server config
	srv := &http.Server{
		Handler:      r,
		Addr:         ":9090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Println("http service running on [::]", srv.Addr)
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(tc)
}
