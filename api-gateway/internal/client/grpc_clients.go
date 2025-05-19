package client

import (
	"log"

	"google.golang.org/grpc"

	carPb "github.com/Akan15/carrental3/car-service/proto"
	rentalPb "github.com/Akan15/carrental3/rental-service/proto"
	userPb "github.com/Akan15/carrental3/user-service/proto"
)

type Clients struct {
	User   userPb.UserServiceClient
	Car    carPb.CarServiceClient
	Rental rentalPb.RentalServiceClient
}

func InitClients() *Clients {
	userConn, err := grpc.Dial("user-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user-service: %v", err)
	}

	carConn, err := grpc.Dial("car-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to car-service: %v", err)
	}

	rentalConn, err := grpc.Dial("rental-service:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to rental-service: %v", err)
	}

	return &Clients{
		User:   userPb.NewUserServiceClient(userConn),
		Car:    carPb.NewCarServiceClient(carConn),
		Rental: rentalPb.NewRentalServiceClient(rentalConn),
	}
}
