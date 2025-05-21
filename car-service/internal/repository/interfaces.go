package repository

import "github.com/Akan15/carrental3/car-service/internal/models"

type CarRepository interface {
	Create(car *models.Car) (*models.Car, error)
	GetByID(id string) (*models.Car, error)
	Update(car *models.Car) error
	Delete(id string) error
	List() ([]*models.Car, error)
	FindByStatus(status string) ([]*models.Car, error)
	FindByCity(city string) ([]*models.Car, error)
	FindByModel(model string) ([]*models.Car, error)
	FindNearby(lat, lon, radius float64) ([]*models.Car, error)
	ChangeStatus(id, status string) (*models.Car, error)
	GetLocation(id string) (float64, float64, error)
	UpdateLocation(id string, lat, lon float64) (*models.Car, error)
}
