package handlers

import (
	"context"

	"github.com/Akan15/carrental3/user-service/internal/metrics"
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
	metrics.RequestCount.WithLabelValues("Register", "POST").Inc()

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

func (h *UserHandler) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := h.usecase.ListUsers()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list users: %v", err)
	}

	var pbUsers []*pb.GetUserResponse
	for _, u := range users {
		pbUsers = append(pbUsers, &pb.GetUserResponse{
			Id:    u.ID.Hex(),
			Name:  u.Name,
			Email: u.Email,
		})
	}

	return &pb.ListUsersResponse{Users: pbUsers}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	err := h.usecase.UpdateUser(req.Id, req.Name, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update user: %v", err)
	}
	return &pb.UpdateUserResponse{
		Message: "User updated successfully",
	}, nil
}

func (h *UserHandler) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	err := h.usecase.UpdatePassword(req.Id, req.OldPassword, req.NewPassword)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "update password failed: %v", err)
	}
	return &pb.UpdatePasswordResponse{Message: "Password updated"}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := h.usecase.DeleteUser(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "delete failed: %v", err)
	}
	return &pb.DeleteUserResponse{Message: "User deleted"}, nil
}

func (h *UserHandler) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequest) (*pb.UpdateEmailResponse, error) {
	err := h.usecase.UpdateEmail(req.Id, req.Password, req.NewEmail)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "email update failed: %v", err)
	}
	return &pb.UpdateEmailResponse{Message: "Email updated successfully"}, nil
}

func (h *UserHandler) FindByName(ctx context.Context, req *pb.FindByNameRequest) (*pb.FindByNameResponse, error) {
	users, err := h.usecase.FindByName(req.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "search failed: %v", err)
	}
	var result []*pb.GetUserResponse
	for _, u := range users {
		result = append(result, &pb.GetUserResponse{
			Id:    u.ID.Hex(),
			Name:  u.Name,
			Email: u.Email,
		})
	}
	return &pb.FindByNameResponse{Users: result}, nil
}

func (h *UserHandler) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	err := h.usecase.ChangePassword(req.Email, req.OldPassword, req.NewPassword)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "password change failed: %v", err)
	}
	return &pb.ChangePasswordResponse{Message: "Password changed"}, nil
}

func (h *UserHandler) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	newToken, err := h.usecase.RefreshToken(req.Token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token refresh failed: %v", err)
	}
	return &pb.RefreshTokenResponse{Token: newToken}, nil
}

func (h *UserHandler) ResendVerification(ctx context.Context, req *pb.ResendVerificationRequest) (*pb.ResendVerificationResponse, error) {
	err := h.usecase.ResendVerification(req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "resend failed: %v", err)
	}
	return &pb.ResendVerificationResponse{Message: "Verification email sent"}, nil
}
