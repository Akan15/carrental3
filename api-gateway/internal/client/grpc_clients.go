package client

import (
	"log"

	"google.golang.org/grpc"

	carpb "carrental/api-gateway/proto/car"
	rentalpb "carrental/api-gateway/proto/rental"
	userpb "carrental/api-gateway/proto/user"
)

type Clients struct {
	UserClient   userpb.UserServiceClient
	CarClient    carpb.CarServiceClient
	RentalClient rentalpb.RentalServiceClient
}

func NewClients() *Clients {
	userConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to user service: %v", err)
	}

	carConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to car service: %v", err)
	}

	rentalConn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to rental service: %v", err)
	}

	return &Clients{
		UserClient:   userpb.NewUserServiceClient(userConn),
		CarClient:    carpb.NewCarServiceClient(carConn),
		RentalClient: rentalpb.NewRentalServiceClient(rentalConn),
	}
}
