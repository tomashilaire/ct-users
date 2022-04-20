package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"test/internal/core/services/testsrv"
	"test/internal/handlers/testhttphdl"
	"test/internal/repositories/testmongorepo"
	"test/pkg/uidgen"
	"time"

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

	tr := testmongorepo.NewTestRepository(db)
	ts := testsrv.NewService(tr, uidgen.New())
	th := testhttphdl.NewHTTPHandler(ts)

	r := mux.NewRouter()

	getRouter := r.Methods("GET").Subrouter()
	getRouter.HandleFunc("/tests", th.GetAllTests)
	getRouter.HandleFunc("/tests/{id:[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}}", th.GetTest)

	postRouter := r.Methods("POST").Subrouter()
	postRouter.HandleFunc("/tests", th.PostTest)

	putRouter := r.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/tests/{id:[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}}", th.PutTest)

	deleteRouter := r.Methods("DELET").Subrouter()
	deleteRouter.HandleFunc("/tests/{id:[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}}", th.DeleteTest)

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
