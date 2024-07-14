package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRounting() {
	r := mux.NewRouter()

	r.HandleFunc("/create", CreateEmployee).Methods("POST")

	r.HandleFunc("/add", GetEmployees).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}