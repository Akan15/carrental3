package usecase

import (
	"github.com/Akan15/carrental3/car-service/internal/models"
	"github.com/Akan15/carrental3/car-service/internal/repository"
)

type CarUseCase struct {
	repo repository.CarRepository
}

func NewCarUseCase(r repository.CarRepository) *CarUseCase {
	return &CarUseCase{repo: r}
}

func (u *CarUseCase) GetCarByID(id string) (*models.Car, error) {
	return u.repo.GetCarByID(id)
}
