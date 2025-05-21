package repository

import (
	"time"

	"github.com/Akan15/carrental3/rental-service/internal/models"
)

type RentalRepositoryInterface interface {
	Create(r *models.Rental) (*models.Rental, error)
	End(id string, end time.Time, price float64) (*models.Rental, error)
	GetByID(id string) (*models.Rental, error)
	List() ([]*models.Rental, error)
	ListByUser(userID string) ([]*models.Rental, error)
	ListByCar(carID string) ([]*models.Rental, error)
	GetActive(userID, carID string) (*models.Rental, error)
	GetActiveRentals() ([]*models.Rental, error)
	Delete(id string) error
	UpdateType(id, rentalType string) (*models.Rental, error)
}
