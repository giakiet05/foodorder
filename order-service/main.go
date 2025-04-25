package main

import (
	"fmt"
	"github.com/giakiet05/foodorder/order-service/client"
	"github.com/giakiet05/foodorder/order-service/db"
	"github.com/giakiet05/foodorder/order-service/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db.Init()

	r := mux.NewRouter()
	r.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
	r.HandleFunc("/orders", handlers.GetOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", handlers.GetOrder).Methods("GET")

	client.InitUserClient() // Khởi tạo kết nối gRPC tới user-service

	exists := client.CheckUserExists(1)
	fmt.Println("User exists:", exists)

	log.Println("✅ Order service running on port 8002")
	log.Fatal(http.ListenAndServe(":8002", r))
}
