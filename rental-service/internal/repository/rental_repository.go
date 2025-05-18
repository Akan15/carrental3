package repository

import (
	"github.com/Akan15/carrental3/rental-service/internal/models"
)

type RentalRepository interface {
	GetRentalByID(id string) (*models.Rental, error)
}

type InMemoryRentalRepo struct {
	data map[string]models.Rental
}

func NewInMemoryRentalRepo() *InMemoryRentalRepo {
	return &InMemoryRentalRepo{
		data: map[string]models.Rental{
			"1": {ID: "1", UserID: "1", CarID: "2", StartTime: "2025-05-13T10:00", EndTime: "2025-05-13T12:00"},
		},
	}
}

func (r *InMemoryRentalRepo) GetRentalByID(id string) (*models.Rental, error) {
	rental, exists := r.data[id]
	if !exists {
		return nil, nil
	}
	return &rental, nil
}
