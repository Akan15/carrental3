package main

import (
	"log"
	"net"

	"github.com/Akan15/carrental3/rental-service/internal/handlers"
	"github.com/Akan15/carrental3/rental-service/internal/repository"
	"github.com/Akan15/carrental3/rental-service/internal/usecase"
	pb "github.com/Akan15/carrental3/rental-service/proto"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func InitMongo() *mongo.Database {
	uri := "mongodb://admin:secret@mongo:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}

	log.Println("✅ Connected to MongoDB")
	return client.Database("carrental")
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	db := InitMongo()
	repo := repository.NewRentalRepository(db)
	uc := usecase.NewRentalUseCase(repo)
	handler := handlers.NewRentalHandler(uc)

	grpcServer := grpc.NewServer()
	pb.RegisterRentalServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	log.Println("✅ RentalService is running on port :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
