package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	handlers "github.com/pattadob/learn-demo-api-with-golang/handlers"
)

func main() {
	log.Print("Starting the service")

	port := os.Getenv("APIPORT")

	if port == "" {
		log.Fatal("Post is not set.")
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	router := handlers.Router()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	log.Print("The service is ready to listen and serve.")

	killSignal := <-stop
	switch killSignal {
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("GOT: SIGTERM...")
	}

	log.Print("The service is shotting down...")
	srv.Shutdown(context.Background())
	log.Print("Server gracefully stopped")
}
