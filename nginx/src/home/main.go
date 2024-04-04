package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the home page!"))
}

func main() {
	log.Println("Starting home service.")

	router := mux.NewRouter()
	homeRouter := router.PathPrefix("/home/api/v1").Subrouter()

	homeRouter.HandleFunc("/welcome", WelcomePage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", homeRouter))
}
