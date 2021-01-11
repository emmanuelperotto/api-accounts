package main

import (
	"accounts/infra"
	"accounts/middlewares"
	"accounts/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	//SETUP DB
	if err := infra.SetupDB(); err != nil {
		log.Fatal("Error setting the DB up => ", err)
	}

	//SETUP ROUTES
	router := mux.NewRouter()
	router.Use(middlewares.Logging)

	router.HandleFunc("/accounts", routes.CreateAccount).Methods("POST")

	//START SERVER
	address := ":3000"
	server := http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("Running server on port ", address)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error running server: ", err)
	}
}
