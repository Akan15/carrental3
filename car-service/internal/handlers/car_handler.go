package handlers

import (
	"context"

	"github.com/Akan15/carrental3/car-service/internal/models"
	"github.com/Akan15/carrental3/car-service/internal/usecase"
	pb "github.com/Akan15/carrental3/car-service/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CarHandler struct {
	pb.UnimplementedCarServiceServer
	usecase *usecase.CarUseCase
}

func NewCarHandler(u *usecase.CarUseCase) *CarHandler {
	return &CarHandler{usecase: u}
}

func (h *CarHandler) CreateCar(ctx context.Context, req *pb.CreateCarRequest) (*pb.Car, error) {
	car := &models.Car{
		Brand:     req.Brand,
		Model:     req.Model,
		City:      req.City,
		Status:    req.Status,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
	created, err := h.usecase.Create(car)
	if err != nil {
		return nil, err
	}
	return mapCarToProto(created), nil
}

func (h *CarHandler) GetCar(ctx context.Context, req *pb.CarIdRequest) (*pb.Car, error) {
	car, err := h.usecase.GetByID(req.Id)
	if err != nil {
		return nil, err
	}
	return mapCarToProto(car), nil
}

func (h *CarHandler) UpdateCar(ctx context.Context, req *pb.UpdateCarRequest) (*pb.Car, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	car := &models.Car{
		ID:        objID,
		Brand:     req.Brand,
		Model:     req.Model,
		City:      req.City,
		Status:    req.Status,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
	if err := h.usecase.Update(car); err != nil {
		return nil, err
	}
	return mapCarToProto(car), nil
}

func (h *CarHandler) DeleteCar(ctx context.Context, req *pb.CarIdRequest) (*pb.Empty, error) {
	if err := h.usecase.Delete(req.Id); err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (h *CarHandler) ListCars(ctx context.Context, _ *pb.Empty) (*pb.CarList, error) {
	cars, err := h.usecase.List()
	if err != nil {
		return nil, err
	}
	var protoCars []*pb.Car
	for _, c := range cars {
		protoCars = append(protoCars, mapCarToProto(c))
	}
	return &pb.CarList{Cars: protoCars}, nil
}

func mapCarToProto(car *models.Car) *pb.Car {
	return &pb.Car{
		Id:        car.ID.Hex(),
		Brand:     car.Brand,
		Model:     car.Model,
		City:      car.City,
		Status:    car.Status,
		Latitude:  car.Latitude,
		Longitude: car.Longitude,
	}
}
