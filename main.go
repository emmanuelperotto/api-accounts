package main

import (
	"accounts/infra"
	"accounts/middlewares"
	"accounts/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func main() {
	//Initializing dotenv
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	//SETUP DB
	if err := infra.SetupDB(); err != nil {
		log.Fatal("[DB Setup Error]", err)
	}

	//SETUP AWS-SDK
	infra.SetupMessagingService()

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
