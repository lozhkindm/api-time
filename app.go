package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func start() {
	router := mux.NewRouter()

	router.HandleFunc("/api/time", timeHandler)

	log.Fatal(http.ListenAndServe("localhost:5555", router))
}
