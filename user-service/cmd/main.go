package main

import (
	"log"
	"net"

	"user-service/internal/handlers"
	"user-service/internal/repository"
	"user-service/internal/usecase"
	pb "user-service/proto"

	"google.golang.org/grpc"
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

	log.Println("UserService is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
