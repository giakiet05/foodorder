package client

import (
	"google.golang.org/grpc"
	"log"
)

var userClient userpb.UserServiceClient

func InitUserClient() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Cannot connect to use-service gRPC", err)
	}

	userClient

}
