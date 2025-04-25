package main

import (
	"github.com/giakiet05/foodorder/user-service/db"
	grpcserver "github.com/giakiet05/foodorder/user-service/grpc"
	"github.com/giakiet05/foodorder/user-service/handlers"
	"github.com/giakiet05/foodorder/user-service/proto/userpb"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	db.Init()
	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUserById).Methods("GET")

	go func() {
		lis, _ := net.Listen("tcp", ":50051")
		s := grpc.NewServer()
		userpb.RegisterUserServiceServer(s, &grpcserver.Server{})
		log.Println("gRPC server on port 50051")
		s.Serve(lis)
	}()

	log.Println("User service running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
