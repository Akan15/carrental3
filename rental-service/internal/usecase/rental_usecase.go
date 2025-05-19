package usecase

import (
	"errors"
	"time"

	"github.com/Akan15/carrental3/rental-service/internal/models"
	"github.com/Akan15/carrental3/rental-service/internal/repository"
)

type RentalUseCase struct {
	repo *repository.RentalRepository
}

func NewRentalUseCase(r *repository.RentalRepository) *RentalUseCase {
	return &RentalUseCase{repo: r}
}

func (u *RentalUseCase) Create(userID, carID, rentalType string) (*models.Rental, error) {
	rental := &models.Rental{
		UserID: userID,
		CarID:  carID,
		Type:   rentalType,
	}
	return u.repo.Create(rental)
}

func (u *RentalUseCase) End(id string) (*models.Rental, error) {
	rental, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if rental.EndTime != nil {
		return nil, errors.New("rental already ended")
	}

	endTime := time.Now()
	duration := endTime.Sub(rental.StartTime)

	var price float64
	switch rental.Type {
	case "normal":
		price = 2000
	case "per_minute":
		minutes := int(duration.Minutes())
		price = float64(minutes) * 30
	default:
		return nil, errors.New("invalid rental type")
	}

	return u.repo.End(id, endTime, price)
}

func (u *RentalUseCase) Get(id string) (*models.Rental, error) {
	return u.repo.GetByID(id)
}

func (u *RentalUseCase) List() ([]*models.Rental, error) {
	return u.repo.List()
}
