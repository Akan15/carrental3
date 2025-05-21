package main

import (
	"log"
	"net"

	"github.com/Akan15/carrental3/car-service/internal/handlers"
	"github.com/Akan15/carrental3/car-service/internal/nats"
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

	db := repository.InitMongo()
	repo := repository.NewCarRepository(db)
	uc := usecase.NewCarUseCase(repo)

	handler := handlers.NewCarHandler(uc)

	grpcServer := grpc.NewServer()
	pb.RegisterCarServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	log.Println("✅ CarService is running on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	go func() {
		err := nats.SubscribeToRentalCreated("nats://localhost:4222", uc)
		if err != nil {
			log.Fatal("NATS subscription error:", err)
		}
	}()

}
