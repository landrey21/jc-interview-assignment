package hash

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/landrey21/jc-interview-assignment/lib/auth"
)

const (
	RESPONSE_DELAY = 5
)

type HashHandler struct{}
type StatsHandler struct{}

type Stats struct {
	Total     int64 `json:"total"`
	Average   int64 `json:"average"`
	TotalTime int64 `json:"-"`
}

var statsData = Stats{0, 0, 0} // stats tracker

// --------------------------------------------------------
func (s *StatsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("/stats")
	sd, err := json.Marshal(statsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(sd)
}

// --------------------------------------------------------
func (h *HashHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("/hash")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	password := r.Form.Get("password")
	if password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing required field 'password'"))
		return
	}

	start := time.Now()

	encodedPassword := auth.HashEncode(password)

	end := time.Since(start)

	statsData.Total++
	//statsData.TotalTime += int64(end) // nanoseconds
	statsData.TotalTime += end.Nanoseconds() / 1e6 // converts to milliseconds
	statsData.Average = (statsData.TotalTime / statsData.Total)

	// wait for some number of seconds before responding, part of the assignment
	time.Sleep(RESPONSE_DELAY * time.Second)

	w.Write([]byte(encodedPassword))
}
