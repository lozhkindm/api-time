package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	loc, err := time.LoadLocation("UTC")

	if err != nil {
		log.Fatal(err)
	}

	tr := timeResource{
		CurrentTime: time.Now().In(loc),
	}

	err = json.NewEncoder(w).Encode(tr)

	if err != nil {
		log.Fatal(err)
	}
}
