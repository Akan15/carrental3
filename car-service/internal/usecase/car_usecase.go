package usecase

import (
	"github.com/Akan15/carrental3/car-service/internal/models"
	"github.com/Akan15/carrental3/car-service/internal/repository"
)

type CarUseCase struct {
	repo *repository.CarRepository
}

func NewCarUseCase(r *repository.CarRepository) *CarUseCase {
	return &CarUseCase{repo: r}
}

func (u *CarUseCase) Create(car *models.Car) (*models.Car, error) {
	return u.repo.Create(car)
}

func (u *CarUseCase) GetByID(id string) (*models.Car, error) {
	return u.repo.GetByID(id)
}

func (u *CarUseCase) Update(car *models.Car) error {
	return u.repo.Update(car)
}

func (u *CarUseCase) Delete(id string) error {
	return u.repo.Delete(id)
}

func (u *CarUseCase) List() ([]*models.Car, error) {
	return u.repo.List()
}
