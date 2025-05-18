package usecase

import (
	"github.com/Akan15/carrental3/rental-service/internal/models"
	"github.com/Akan15/carrental3/rental-service/internal/repository"
)

type RentalUseCase struct {
	repo repository.RentalRepository
}

func NewRentalUseCase(r repository.RentalRepository) *RentalUseCase {
	return &RentalUseCase{repo: r}
}

func (u *RentalUseCase) GetRentalByID(id string) (*models.Rental, error) {
	return u.repo.GetRentalByID(id)
}
