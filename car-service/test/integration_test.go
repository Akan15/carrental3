package test

import (
	"context"
	"log"
	"os"
	"testing"

	pb "github.com/Akan15/carrental3/car-service/proto"
	"google.golang.org/grpc"
)

var client pb.CarServiceClient

func TestMain(m *testing.M) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ failed to connect: %v", err)
	}
	client = pb.NewCarServiceClient(conn)
	os.Exit(m.Run())
}

func TestCreateAndGetCar(t *testing.T) {
	resp, err := client.CreateCar(context.Background(), &pb.CreateCarRequest{
		Brand:     "Hyundai",
		Model:     "Sonata",
		City:      "Almaty",
		Status:    "free",
		Latitude:  43.23,
		Longitude: 76.88,
	})
	if err != nil {
		t.Fatalf("❌ CreateCar failed: %v", err)
	}

	car, err := client.GetCar(context.Background(), &pb.CarIdRequest{Id: resp.Id})
	if err != nil {
		t.Fatalf("❌ GetCar failed: %v", err)
	}
	if car.Brand != "Hyundai" {
		t.Errorf("Expected Hyundai, got %s", car.Brand)
	}
}
