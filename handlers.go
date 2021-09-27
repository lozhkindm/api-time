package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tz := r.URL.Query().Get("tz")

	if tz == "" {
		tz = UTC
	}

	w.Header().Add("Content-Type", "application/json")

	loc, err := time.LoadLocation(tz)

	if err != nil {
		http.Error(w, "invalid timezone", http.StatusNotFound)
		return
	}

	tr := timeResource{
		CurrentTime: time.Now().In(loc),
	}

	err = json.NewEncoder(w).Encode(tr)

	if err != nil {
		log.Fatal(err)
	}
}
