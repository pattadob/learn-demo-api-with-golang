package main

import (
	"log"
	"net/http"
	"os"

	handlers "github.com/pattadob/learn-demo-api-with-golang/handlers"
)

func main() {
	log.Print("Starting the service")

	port := os.Getenv("APIPORT")

	if port == "" {
		log.Fatal("Post is not set.")
	}

	router := handlers.Router()

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, router))

}
