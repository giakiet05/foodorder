package main

import (
	"github.com/giakiet05/foodorder/food-service/db"
	"github.com/giakiet05/foodorder/food-service/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db.Init()

	r := mux.NewRouter()

	r.HandleFunc("/foods", handlers.CreateFood).Methods("POST")
	r.HandleFunc("/foods", handlers.GetFoods).Methods("GET")
	r.HandleFunc("/foods/{id}", handlers.GetFood).Methods("GET")

	log.Println("Food service running on port 8001")
	log.Fatalln(http.ListenAndServe(":8001", r))
}
