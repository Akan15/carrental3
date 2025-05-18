package main

import (
	"log"
	"net"

	"github.com/Akan15/carrental3/car-service/internal/handlers"
	"github.com/Akan15/carrental3/car-service/internal/repository"
	"github.com/Akan15/carrental3/car-service/internal/usecase"
	pb "github.com/Akan15/carrental3/car-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	repo := repository.NewInMemoryCarRepo()
	uc := usecase.NewCarUseCase(repo)
	handler := handlers.NewCarHandler(uc)

	grpcServer := grpc.NewServer()
	pb.RegisterCarServiceServer(grpcServer, handler)

	// Enable reflection
	reflection.Register(grpcServer)

	log.Println("CarService is running on port :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
