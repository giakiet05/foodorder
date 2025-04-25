package grpcserver

import (
	"context"
	"github.com/giakiet05/foodorder/food-service/db"
	"github.com/giakiet05/foodorder/food-service/models"
	"github.com/giakiet05/foodorder/food-service/proto/foodpb"
)

// Server là struct triển khai service gRPC
type Server struct {
	foodpb.UnimplementedFoodServiceServer
}

// CheckFood kiểm tra món ăn có tồn tại không
func (s *Server) CheckFood(ctx context.Context, req *foodpb.FoodRequest) (*foodpb.FoodResponse, error) {
	var food models.Food
	if err := db.DB.First(&food, req.Id).Error; err != nil {
		// Không tìm thấy
		return &foodpb.FoodResponse{Exists: false}, nil
	}
	return &foodpb.FoodResponse{Exists: true}, nil
}
