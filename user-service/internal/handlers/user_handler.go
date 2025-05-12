package handlers

import (
	"context"
	"user-service/internal/usecase"
	pb "user-service/proto"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	usecase *usecase.UserUseCase
}

func NewUserHandler(u *usecase.UserUseCase) *UserHandler {
	return &UserHandler{usecase: u}
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.usecase.GetUserByID(req.Id)
	if err != nil || user == nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
