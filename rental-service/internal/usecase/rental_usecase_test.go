package usecase_test

import (
	"testing"
	"time"

	"github.com/Akan15/carrental3/rental-service/internal/models"
	"github.com/Akan15/carrental3/rental-service/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// 🔧 Мок для RentalRepositoryInterface
type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) Create(r *models.Rental) (*models.Rental, error) {
	args := m.Called(r)
	return args.Get(0).(*models.Rental), args.Error(1)
}

func (m *mockRepo) End(id string, end time.Time, price float64) (*models.Rental, error) {
	args := m.Called(id, end, price)
	return args.Get(0).(*models.Rental), args.Error(1)
}

func (m *mockRepo) GetByID(id string) (*models.Rental, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Rental), args.Error(1)
}

func (m *mockRepo) List() ([]*models.Rental, error) {
	args := m.Called()
	return args.Get(0).([]*models.Rental), args.Error(1)
}
func (m *mockRepo) ListByUser(userID string) ([]*models.Rental, error) {
	args := m.Called(userID)
	return args.Get(0).([]*models.Rental), args.Error(1)
}

func (m *mockRepo) ListByCar(carID string) ([]*models.Rental, error) {
	args := m.Called(carID)
	return args.Get(0).([]*models.Rental), args.Error(1)
}

func (m *mockRepo) GetActive(userID, carID string) (*models.Rental, error) {
	args := m.Called(userID, carID)
	return args.Get(0).(*models.Rental), args.Error(1)
}

func (m *mockRepo) GetActiveRentals() ([]*models.Rental, error) {
	args := m.Called()
	return args.Get(0).([]*models.Rental), args.Error(1)
}

func (m *mockRepo) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *mockRepo) UpdateType(id, rentalType string) (*models.Rental, error) {
	args := m.Called(id, rentalType)
	return args.Get(0).(*models.Rental), args.Error(1)
}

func TestCreateRental(t *testing.T) {
	repo := new(mockRepo)
	uc := usecase.NewRentalUseCase(repo, nil)

	input := &models.Rental{UserID: "u1", CarID: "c1", Type: "normal"}
	repo.On("Create", mock.AnythingOfType("*models.Rental")).Return(input, nil)

	result, err := uc.Create("u1", "c1", "normal")

	assert.NoError(t, err)
	assert.Equal(t, "u1", result.UserID)
	assert.Equal(t, "c1", result.CarID)
	assert.Equal(t, "normal", result.Type)
}

func TestEndRental(t *testing.T) {
	repo := new(mockRepo)
	uc := usecase.NewRentalUseCase(repo, nil)

	rental := &models.Rental{
		ID:        [12]byte{},
		UserID:    "u1",
		CarID:     "c1",
		Type:      "per_minute",
		StartTime: time.Now().Add(-10 * time.Minute),
		EndTime:   nil,
	}

	repo.On("GetByID", "1").Return(rental, nil)
	repo.On("End", "1", mock.AnythingOfType("time.Time"), 300.0).Return(rental, nil)

	result, err := uc.End("1")

	assert.NoError(t, err)
	assert.Equal(t, "u1", result.UserID)
}

func TestEndRentalAlreadyEnded(t *testing.T) {
	repo := new(mockRepo)
	uc := usecase.NewRentalUseCase(repo, nil)

	now := time.Now()
	rental := &models.Rental{
		ID:      [12]byte{},
		UserID:  "u1",
		CarID:   "c1",
		Type:    "normal",
		EndTime: &now,
	}

	repo.On("GetByID", "1").Return(rental, nil)

	_, err := uc.End("1")
	assert.Error(t, err)
	assert.Equal(t, "rental already ended", err.Error())
}

func TestEndRentalInvalidType(t *testing.T) {
	repo := new(mockRepo)
	uc := usecase.NewRentalUseCase(repo, nil)

	rental := &models.Rental{
		ID:        [12]byte{},
		UserID:    "u1",
		CarID:     "c1",
		Type:      "invalid_type",
		StartTime: time.Now(),
		EndTime:   nil,
	}

	repo.On("GetByID", "1").Return(rental, nil)

	_, err := uc.End("1")
	assert.Error(t, err)
	assert.Equal(t, "invalid rental type", err.Error())
}
