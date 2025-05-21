package usecase

import (
	"errors"
	"time"

	"github.com/Akan15/carrental3/rental-service/internal/models"
	"github.com/Akan15/carrental3/rental-service/internal/nats"
	"github.com/Akan15/carrental3/rental-service/internal/repository"
)

type RentalUseCase struct {
	repo          repository.RentalRepositoryInterface
	NatsPublisher *nats.Publisher
}

func NewRentalUseCase(repo repository.RentalRepositoryInterface, publisher *nats.Publisher) *RentalUseCase {
	return &RentalUseCase{
		repo:          repo,
		NatsPublisher: publisher,
	}
}

func (u *RentalUseCase) Create(userID, carID, rentalType string) (*models.Rental, error) {
	rental := &models.Rental{
		UserID: userID,
		CarID:  carID,
		Type:   rentalType,
	}
	r, err := u.repo.Create(rental)
	if err != nil {
		return nil, err
	}

	// Публикация события
	if u.NatsPublisher != nil {
		_ = u.NatsPublisher.PublishRentalCreated(nats.RentalCreatedEvent{
			RentalID: r.ID.Hex(),
			CarID:    carID,
			Status:   "occupied",
		})
	}

	return r, nil
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

func (u *RentalUseCase) ListByUser(userID string) ([]*models.Rental, error) {
	return u.repo.ListByUser(userID)
}

func (u *RentalUseCase) ListByCar(carID string) ([]*models.Rental, error) {
	return u.repo.ListByCar(carID)
}

func (u *RentalUseCase) GetActive(userID, carID string) (*models.Rental, error) {
	return u.repo.GetActive(userID, carID)
}

func (u *RentalUseCase) Delete(id string) error {
	return u.repo.Delete(id)
}

func (u *RentalUseCase) UpdateType(id, rentalType string) (*models.Rental, error) {
	return u.repo.UpdateType(id, rentalType)
}

func (u *RentalUseCase) GetActiveRentals() ([]*models.Rental, error) {
	return u.repo.GetActiveRentals()
}
