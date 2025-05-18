package main

import (
	"log"
	"net"

	"github.com/Akan15/carrental3/rental-service/internal/handlers"
	"github.com/Akan15/carrental3/rental-service/internal/repository"
	"github.com/Akan15/carrental3/rental-service/internal/usecase"
	pb "github.com/Akan15/carrental3/rental-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	repo := repository.NewInMemoryRentalRepo()
	uc := usecase.NewRentalUseCase(repo)
	handler := handlers.NewRentalHandler(uc)

	grpcServer := grpc.NewServer()
	pb.RegisterRentalServiceServer(grpcServer, handler)

	// ✅ Добавь это, чтобы grpcurl работал
	reflection.Register(grpcServer)

	log.Println("RentalService running on :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
