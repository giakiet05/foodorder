package client

import (
	"context"
	"github.com/giakiet05/foodorder/food-service/proto/foodpb"
	"google.golang.org/grpc"
	"log"
	"time"
)

var foodClient foodpb.FoodServiceClient

func InitFoodClient() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatal("❌ Cannot connect to food gRPC:", err)
	}
	foodClient = foodpb.NewFoodServiceClient(conn)
}

func CheckFoodExists(foodID uint) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := foodClient.CheckFood(ctx, &foodpb.FoodRequest{Id: uint32(foodID)})
	if err != nil {
		log.Println("gRPC error (food):", err)
		return false
	}
	return resp.Exists
}
