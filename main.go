package main

import (
	"accounts/entities"
	"accounts/infra"
	"accounts/middlewares"
	"accounts/routes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
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

	//FIXME: extract to infra package
	//SETUP AWS-SDK
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           os.Getenv("AWS_PROFILE"),
		SharedConfigState: session.SharedConfigEnable,
	}))

	snsClient := sns.New(sess)

	acc := entities.Account{
		ID:          10,
		Code:        "41325",
		Agency:      "0001",
		AccessToken: "",
	}
	input := sns.PublishInput{
		Message: aws.String("AccountCreated"),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"ID": {
				DataType:    aws.String("String"),
				StringValue: aws.String(strconv.FormatInt(acc.ID, 10)),
			},
			"Code": {
				DataType:    aws.String("String"),
				StringValue: aws.String(acc.Code),
			},
			"Agency": {
				DataType:    aws.String("String"),
				StringValue: aws.String(acc.Agency),
			},
		},
		TopicArn: aws.String(os.Getenv("TOPIC_ARN")),
	}
	output, err := snsClient.Publish(&input)
	if err != nil {
		log.Println("[Publish error]", err)
	}
	fmt.Println(output)

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
