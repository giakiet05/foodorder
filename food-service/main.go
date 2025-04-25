package main

import (
	"github.com/giakiet05/foodorder/food-service/db"
	grpcserver "github.com/giakiet05/foodorder/food-service/grpc"
	"github.com/giakiet05/foodorder/food-service/handlers"
	"github.com/giakiet05/foodorder/food-service/proto/foodpb"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	db.Init()

	r := mux.NewRouter()

	go func() {
		lis, _ := net.Listen("tcp", ":50052") // Dùng cổng khác với user-service
		s := grpc.NewServer()
		foodpb.RegisterFoodServiceServer(s, &grpcserver.Server{})
		log.Println("🔌 gRPC Food service running on port 50052")
		s.Serve(lis)
	}()

	r.HandleFunc("/foods", handlers.CreateFood).Methods("POST")
	r.HandleFunc("/foods", handlers.GetFoods).Methods("GET")
	r.HandleFunc("/foods/{id}", handlers.GetFood).Methods("GET")

	log.Println("Food service running on port 8001")
	log.Fatalln(http.ListenAndServe(":8001", r))
}
