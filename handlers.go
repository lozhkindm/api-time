package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	tz := getTimezone(r)
	tzs := strings.Split(tz, ",")

	if len(tzs) > 1 {
		multipleTimezoneHandler(w, tzs)
	} else {
		singleTimezoneHandler(w, tz)
	}
}

func multipleTimezoneHandler(w http.ResponseWriter, tzs []string) {
	tzmap := make(map[string]time.Time)

	for _, tz := range tzs {
		loc, err := time.LoadLocation(tz)

		if err != nil {
			http.Error(w, "invalid timezone", http.StatusNotFound)
			return
		}

		tzmap[tz] = time.Now().In(loc)
	}

	err := json.NewEncoder(w).Encode(tzmap)

	if err != nil {
		log.Fatal(err)
	}
}

func singleTimezoneHandler(w http.ResponseWriter, tz string) {
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

func getTimezone(r *http.Request) string {
	tz := r.URL.Query().Get("tz")

	if tz == "" {
		tz = UTC
	}

	return tz
}
