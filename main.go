package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	ENV_PREFIX          = "JC_INTERVIEW_ASSIGNMENT_"
	DEFAULT_ENVIRONMENT = "prod"
)

func main() {

	port := ""
	if len(os.Args) > 1 {
		port = string(os.Args[1])
	} else {
		port = os.Getenv(ENV_PREFIX + "PORT")
	}
	if port == "" {
		log.Fatal("PORT is required as an agrument or ENV var")
	}

	server := &http.Server{Addr: ":" + port}

	connectionsDone := make(chan struct{})

	sig := make(chan os.Signal, 1)

	go func() {
		signal.Notify(sig, os.Interrupt, syscall.SIGINT)
		<-sig

		// We received an interrupt signal, shut down
		if err := server.Shutdown(context.Background()); err != nil {
			// Error closing listeners, or context timeout
			log.Printf("ERROR HTTP server Shutdown: %v", err)
		}
		close(connectionsDone)
	}()

	// Handle the shutdown endpoint here so we can access sig without making it global
	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("/shutdown")
		w.Write([]byte("Shutting Down"))
		sig <- syscall.SIGINT
	})

	// We keep routes/handlers in routes.go
	routes()

	log.Printf("HTTP server starting on PORT %s ...", port)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener
		log.Fatalf("ERROR HTTP server ListenAndServe: %v", err)
	}

	log.Print("HTTP server shutting down ...")
	// Don't let the progam fall through until connections have been closed
	<-connectionsDone
	log.Print("HTTP server done")
}
