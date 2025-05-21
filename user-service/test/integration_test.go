package test

import (
	"context"
	"log"
	"os"
	"testing"

	pb "github.com/Akan15/carrental3/user-service/proto"
	"google.golang.org/grpc"
)

var client pb.UserServiceClient

func TestMain(m *testing.M) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	client = pb.NewUserServiceClient(conn)
	os.Exit(m.Run())
}

func TestRegisterGRPC(t *testing.T) {
	_, err := client.Register(context.Background(), &pb.RegisterRequest{
		Name:     "Test User",
		Email:    "testuser@example.com",
		Password: "testpass",
	})
	if err != nil && err.Error() != "rpc error: code = Internal desc = registration failed: user already exists" {
		t.Fatalf("Register failed: %v", err)
	}
}

func TestLoginGRPC(t *testing.T) {
	resp, err := client.Login(context.Background(), &pb.LoginRequest{
		Email:    "testuser@example.com",
		Password: "testpass",
	})
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}
	if resp.Token == "" {
		t.Errorf("Expected token, got empty")
	}
}

func TestGetUserByID(t *testing.T) {
	// Используем ID из ListUsers (у тебя John — ID: 6829becd4d236e5d60268881)
	userResp, err := client.GetUser(context.Background(), &pb.GetUserRequest{
		Id: "6829becd4d236e5d60268881",
	})

	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}
	if userResp.Email != "john@mail.com" {
		t.Errorf("Expected email john@mail.com, got %v", userResp.Email)
	}
}
