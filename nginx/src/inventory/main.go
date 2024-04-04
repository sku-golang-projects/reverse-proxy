package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InventoryPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the inventory page!"))
}

func main() {
	log.Println("Starting inventory service.")

	router := mux.NewRouter()
	inventoryRouter := router.PathPrefix("/inventory/api/v1").Subrouter()

	inventoryRouter.HandleFunc("/welcome", InventoryPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8082", inventoryRouter))
}
