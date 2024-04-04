package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func OrderPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the order page!"))
}

func main() {
	log.Println("Starting order service.")

	router := mux.NewRouter()
	orderRouter := router.PathPrefix("/order/api/v1").Subrouter()

	orderRouter.HandleFunc("/welcome", OrderPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8083", orderRouter))
}
