package client

import (
	"context"
	"github.com/giakiet05/foodorder/user-service/user-service/proto/userpb"
	"google.golang.org/grpc"
	"log"
	"time"
)

var userClient userpb.UserServiceClient

func InitUserClient() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Cannot connect to use-service gRPC", err)
	}

	userClient = userpb.NewUserServiceClient(conn)

}

func CheckUserExists(userId uint) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := userClient.CheckUser(ctx, &userpb.UserRequest{Id: uint32(userId)})
	if err != nil {
		log.Println("Error calling CheckUser gRPC", err)
		return false
	}
	return res.Exists
}
