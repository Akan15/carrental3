package main

import (
	"log"
	"net"

	"github.com/Akan15/carrental3/user-service/internal/handlers"
	"github.com/Akan15/carrental3/user-service/internal/repository"
	"github.com/Akan15/carrental3/user-service/internal/usecase"
	pb "github.com/Akan15/carrental3/user-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	repo := repository.NewInMemoryUserRepo()
	uc := usecase.NewUserUseCase(repo)
	handler := handlers.NewUserHandler(uc)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, handler)

	// 👉 Включаем reflection
	reflection.Register(grpcServer)

	log.Println("UserService is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
