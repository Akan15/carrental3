package handlers

import (
	"context"

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

func (h *RentalHandler) GetRental(ctx context.Context, req *pb.GetRentalRequest) (*pb.GetRentalResponse, error) {
	rental, err := h.usecase.GetRentalByID(req.Id)
	if err != nil || rental == nil {
		return nil, err
	}

	return &pb.GetRentalResponse{
		Id:        rental.ID,
		UserId:    rental.UserID,
		CarId:     rental.CarID,
		StartTime: rental.StartTime,
		EndTime:   rental.EndTime,
	}, nil
}
