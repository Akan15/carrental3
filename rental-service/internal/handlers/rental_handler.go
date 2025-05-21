package handlers

import (
	"context"
	"time"

	"github.com/Akan15/carrental3/rental-service/internal/models"
	"github.com/Akan15/carrental3/rental-service/internal/usecase"
	pb "github.com/Akan15/carrental3/rental-service/proto"
)

type RentalHandler struct {
	pb.UnimplementedRentalServiceServer
	usecase *usecase.RentalUseCase
}

func NewRentalHandler(u *usecase.RentalUseCase) *RentalHandler {
	return &RentalHandler{usecase: u}
}

func (h *RentalHandler) CreateRental(ctx context.Context, req *pb.CreateRentalRequest) (*pb.Rental, error) {
	rental, err := h.usecase.Create(req.UserId, req.CarId, req.Type)
	if err != nil {
		return nil, err
	}
	return mapToProtoModel(rental), nil
}

func (h *RentalHandler) EndRental(ctx context.Context, req *pb.EndRentalRequest) (*pb.Rental, error) {
	rental, err := h.usecase.End(req.Id)
	if err != nil {
		return nil, err
	}
	return mapToProtoModel(rental), nil
}

func (h *RentalHandler) GetRental(ctx context.Context, req *pb.GetRentalRequest) (*pb.Rental, error) {
	rental, err := h.usecase.Get(req.Id)
	if err != nil {
		return nil, err
	}
	return mapToProtoModel(rental), nil
}

func (h *RentalHandler) ListRentals(ctx context.Context, _ *pb.Empty) (*pb.RentalList, error) {
	rentals, err := h.usecase.List()
	if err != nil {
		return nil, err
	}
	var protoRentals []*pb.Rental
	for _, r := range rentals {
		protoRentals = append(protoRentals, mapToProtoModel(r))
	}
	return &pb.RentalList{Rentals: protoRentals}, nil
}

func mapToProtoModel(m *models.Rental) *pb.Rental {
	var endTime string
	if m.EndTime != nil {
		endTime = m.EndTime.Format(time.RFC3339)
	}
	return &pb.Rental{
		Id:         m.ID.Hex(),
		UserId:     m.UserID,
		CarId:      m.CarID,
		Type:       m.Type,
		StartTime:  m.StartTime.Format(time.RFC3339),
		EndTime:    endTime,
		TotalPrice: m.TotalPrice,
	}
}

func (h *RentalHandler) ListByUser(ctx context.Context, req *pb.UserIdRequest) (*pb.RentalList, error) {
	rentals, err := h.usecase.ListByUser(req.UserId)
	if err != nil {
		return nil, err
	}
	var list []*pb.Rental
	for _, r := range rentals {
		list = append(list, mapToProtoModel(r))
	}
	return &pb.RentalList{Rentals: list}, nil
}

func (h *RentalHandler) ListByCar(ctx context.Context, req *pb.CarIdRequest) (*pb.RentalList, error) {
	rentals, err := h.usecase.ListByCar(req.CarId)
	if err != nil {
		return nil, err
	}
	var list []*pb.Rental
	for _, r := range rentals {
		list = append(list, mapToProtoModel(r))
	}
	return &pb.RentalList{Rentals: list}, nil
}

func (h *RentalHandler) GetActiveRentals(ctx context.Context, _ *pb.Empty) (*pb.RentalList, error) {
	rentals, err := h.usecase.GetActiveRentals()
	if err != nil {
		return nil, err
	}
	var list []*pb.Rental
	for _, r := range rentals {
		list = append(list, mapToProtoModel(r))
	}
	return &pb.RentalList{Rentals: list}, nil
}

func (h *RentalHandler) DeleteRental(ctx context.Context, req *pb.RentalIdRequest) (*pb.Empty, error) {
	err := h.usecase.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (h *RentalHandler) UpdateType(ctx context.Context, req *pb.UpdateTypeRequest) (*pb.Rental, error) {
	rental, err := h.usecase.UpdateType(req.Id, req.Type)
	if err != nil {
		return nil, err
	}
	return mapToProtoModel(rental), nil
}
