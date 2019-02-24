package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	//"time"
	//"github.com/landrey21/jc-interview-assignment/lib/auth"
	//"github.com/landrey21/jc-interview-assignment/routes/hash"
)

const (
	ENV_PREFIX          = "JC_ASSIGNMENT_"
	DEFAULT_ENVIRONMENT = "prod"

//	RESPONSE_DELAY      = 10
)

// type Stats struct {
// 	Total     int64 `json:"total"`
// 	Average   int64 `json:"average"`
// 	TotalTime int64 `json:"-"`
// }

// var statsData = Stats{0, 0, 0} // global stats tracker

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

	// TODO how to handle by GET POST?
	// TODO move to routes
	// http.HandleFunc("/stats", stats)
	// http.HandleFunc("/stats/", stats)
	// http.HandleFunc("/hash", hash)

	// We keep routes/handlers in routes.go
	routes()
	// http.Handle("/stats", &hash.StatsHandler{})
	// http.Handle("/stats/", &hash.StatsHandler{})
	// http.Handle("/hash", &hash.HashHandler{})

	// catch-all route
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write([]byte("PAGE NOT FOUND"))
	// })

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

// // --------------------------------------------------------
// func stats(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("/stats")
// 	s, err := json.Marshal(statsData)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(s)
// }

// // --------------------------------------------------------
// func hash(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("/hash")
// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	password := r.Form.Get("password")
// 	if password == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("Missing required field 'password'"))
// 		return
// 	}

// 	start := time.Now()

// 	encodedPassword := auth.HashEncode(password)

// 	end := time.Since(start)

// 	statsData.Total++
// 	//statsData.TotalTime += int64(end) // nanoseconds
// 	statsData.TotalTime += end.Nanoseconds() / 1e6 // converts to milliseconds
// 	statsData.Average = (statsData.TotalTime / statsData.Total)

// 	// wait for some number of seconds before responding, part of the assignment
// 	time.Sleep(RESPONSE_DELAY * time.Second)

// 	w.Write([]byte(encodedPassword))
// }
