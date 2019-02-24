package main

import (
	"net/http"

	"github.com/landrey21/jc-interview-assignment/routes/hash"
)

func routes() {

	http.Handle("/stats", &hash.StatsHandler{})
	http.Handle("/hash", &hash.HashHandler{})

	// catch-all route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("PAGE NOT FOUND"))
	})
}
