package main

import (
	"accounts/middlewares"
	"accounts/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.Use(middlewares.Logging)

	router.HandleFunc("/accounts", routes.CreateAccount).Methods("POST")

	address := ":3000"
	server := http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("Running server on port ", address)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Error running server: ", err)
	}
}
