package test

import (
	"context"
	"log"
	"testing"

	pb "github.com/Akan15/carrental3/user-service/proto"

	"google.golang.org/grpc"
)

func TestRegisterGRPC(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	resp, err := client.Register(context.Background(), &pb.RegisterRequest{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "123456",
	})
	if err != nil {
		t.Fatalf("Register failed: %v", err)
	}

	log.Println("Register response:", resp.Message)
}
