package grpcserver

import (
	"context"
	"github.com/giakiet05/foodorder/user-service/db"
	"github.com/giakiet05/foodorder/user-service/models"
	"github.com/giakiet05/foodorder/user-service/proto/userpb"
)

type Server struct {
	userpb.UnimplementedUserServiceServer
}

func (s *Server) CheckUser(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	var user models.User
	if err := db.DB.First(&user, req.Id).Error; err != nil {
		return &userpb.UserResponse{Exists: false}, nil
	}
	return &userpb.UserResponse{Exists: true}, nil
}
