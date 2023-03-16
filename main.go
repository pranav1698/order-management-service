package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Extrypoint for the microservice
func main() {
	route := mux.NewRouter()

	route.HandleFunc("/create", CreateNewOrder).Methods("POST")
	route.HandleFunc("/update", UpdateStatus).Methods("PUT")
	route.HandleFunc("/orders", GetOrders).Methods("GET")
	route.HandleFunc("/orders/{filter}/{value}", GetOrder).Methods("GET")

	err := http.ListenAndServe(":8080", route)
	if err != nil {
		log.Fatal(err)
	}
}
