package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/noobg1/hit_men/app/services"

	"github.com/noobg1/hit_men/app/middleware"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/noobg1/hit_men/app/handlers"
)

const (
	WEBSERVERPORT = ":7000"
)

func main() {
	// create connection object
	db, err := services.NewMySQLDatastore()
	if err != nil {
		log.Print(err)
	}
	fmt.Println(db)
	defer db.Close()
	router := mux.NewRouter()
	// routes
	router.HandleFunc("/ping", handlers.Ping)
	router.Handle("/allocate/{username}", handlers.CreateUserAndAllocateTeam(db)).Methods("GET")

	// interceptors
	http.Handle("/", middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, router)))

	http.ListenAndServe(WEBSERVERPORT, nil)

}
