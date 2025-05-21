package usecase

import (
	"context"

	"github.com/Akan15/carrental3/car-service/internal/models"
	"github.com/Akan15/carrental3/car-service/internal/repository"
)

type CarUsecase interface {
	Create(car *models.Car) (*models.Car, error)
	GetByID(id string) (*models.Car, error)
	Update(car *models.Car) error
	Delete(id string) error
	List() ([]*models.Car, error)
	GetAvailableCars() ([]*models.Car, error)
	GetCarsByCity(city string) ([]*models.Car, error)
	GetCarsByStatus(status string) ([]*models.Car, error)
	FindByModel(model string) ([]*models.Car, error)
	FindNearbyCars(lat, lon, radius float64) ([]*models.Car, error)
	ChangeStatus(ctx context.Context, id, status string) error
	GetCarLocation(id string) (float64, float64, error)
	UpdateLocation(id string, lat, lon float64) error
}

type carUsecase struct {
	repo repository.CarRepositoryInterface
}

// ✅ Конструктор возвращает интерфейс
func NewCarUseCase(repo repository.CarRepositoryInterface) CarUsecase {
	return &carUsecase{repo: repo}
}

func (u *carUsecase) Create(car *models.Car) (*models.Car, error) {
	return u.repo.Create(car)
}

func (u *carUsecase) GetByID(id string) (*models.Car, error) {
	return u.repo.GetByID(id)
}

func (u *carUsecase) Update(car *models.Car) error {
	return u.repo.Update(car)
}

func (u *carUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}

func (u *carUsecase) List() ([]*models.Car, error) {
	return u.repo.List()
}

func (u *carUsecase) GetAvailableCars() ([]*models.Car, error) {
	return u.repo.FindByStatus("free")
}

func (u *carUsecase) GetCarsByCity(city string) ([]*models.Car, error) {
	return u.repo.FindByCity(city)
}

func (u *carUsecase) GetCarsByStatus(status string) ([]*models.Car, error) {
	return u.repo.FindByStatus(status)
}

func (u *carUsecase) FindByModel(model string) ([]*models.Car, error) {
	return u.repo.FindByModel(model)
}

func (u *carUsecase) FindNearbyCars(lat, lon, radius float64) ([]*models.Car, error) {
	return u.repo.FindNearby(lat, lon, radius)
}

// 🟢 Исправлено: принимает context.Context
func (u *carUsecase) ChangeStatus(ctx context.Context, id, status string) error {
	return u.repo.ChangeStatus(id, status)
}

func (u *carUsecase) GetCarLocation(id string) (float64, float64, error) {
	return u.repo.GetLocation(id)
}

func (u *carUsecase) UpdateLocation(id string, lat, lon float64) error {
	return u.repo.UpdateLocation(id, lat, lon)
}
