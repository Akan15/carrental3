package handlers

import (
	"context"

	"github.com/Akan15/carrental3/car-service/internal/usecase"
	pb "github.com/Akan15/carrental3/car-service/proto"
)

type CarHandler struct {
	pb.UnimplementedCarServiceServer
	usecase *usecase.CarUseCase
}

func NewCarHandler(u *usecase.CarUseCase) *CarHandler {
	return &CarHandler{usecase: u}
}

func (h *CarHandler) GetCar(ctx context.Context, req *pb.GetCarRequest) (*pb.GetCarResponse, error) {
	car, err := h.usecase.GetCarByID(req.Id)
	if err != nil || car == nil {
		return nil, err
	}
	return &pb.GetCarResponse{
		Id:    car.ID,
		Brand: car.Brand,
		Model: car.Model,
	}, nil
}
