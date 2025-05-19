package handlers

import (
	"context"

	"github.com/Akan15/carrental3/user-service/internal/usecase"
	pb "github.com/Akan15/carrental3/user-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return &pb.GetUserResponse{
		Id:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (h *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	err := h.usecase.Register(req.Name, req.Email, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "registration failed: %v", err)
	}

	return &pb.RegisterResponse{
		Message: "Registration successful",
	}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "login failed: %v", err)
	}

	return &pb.LoginResponse{
		Token: token,
	}, nil
}

func (h *UserHandler) Verify(ctx context.Context, req *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	userID, err := h.usecase.VerifyToken(req.Token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	return &pb.VerifyResponse{
		UserId: userID,
	}, nil
}
