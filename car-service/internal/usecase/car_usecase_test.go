package usecase_test

import (
	"context"
	"testing"

	"github.com/Akan15/carrental3/car-service/internal/models"
	"github.com/Akan15/carrental3/car-service/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) Create(car *models.Car) (*models.Car, error) {
	args := m.Called(car)
	return args.Get(0).(*models.Car), args.Error(1)
}
func (m *mockRepo) GetByID(id string) (*models.Car, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Car), args.Error(1)
}
func (m *mockRepo) Update(car *models.Car) error {
	args := m.Called(car)
	return args.Error(0)
}
func (m *mockRepo) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *mockRepo) List() ([]*models.Car, error) {
	args := m.Called()
	return args.Get(0).([]*models.Car), args.Error(1)
}
func (m *mockRepo) FindByStatus(status string) ([]*models.Car, error) {
	args := m.Called(status)
	return args.Get(0).([]*models.Car), args.Error(1)
}
func (m *mockRepo) FindByCity(city string) ([]*models.Car, error) {
	args := m.Called(city)
	return args.Get(0).([]*models.Car), args.Error(1)
}
func (m *mockRepo) FindByModel(model string) ([]*models.Car, error) {
	args := m.Called(model)
	return args.Get(0).([]*models.Car), args.Error(1)
}
func (m *mockRepo) FindNearby(lat, lon, radius float64) ([]*models.Car, error) {
	args := m.Called(lat, lon, radius)
	return args.Get(0).([]*models.Car), args.Error(1)
}
func (m *mockRepo) ChangeStatus(id, status string) (*models.Car, error) {
	args := m.Called(id, status)
	return args.Get(0).(*models.Car), args.Error(1)
}
func (m *mockRepo) GetLocation(id string) (float64, float64, error) {
	args := m.Called(id)
	return args.Get(0).(float64), args.Get(1).(float64), args.Error(2)
}
func (m *mockRepo) UpdateLocation(id string, lat, lon float64) (*models.Car, error) {
	args := m.Called(id, lat, lon)
	return args.Get(0).(*models.Car), args.Error(1)
}

// 🔹 Тесты:

func TestGetAvailableCars(t *testing.T) {
	m := new(mockRepo)
	expected := []*models.Car{{Brand: "Toyota", Status: "free"}}
	m.On("FindByStatus", "free").Return(expected, nil)

	uc := usecase.NewCarUseCase(m)
	cars, err := uc.GetAvailableCars()

	assert.NoError(t, err)
	assert.Len(t, cars, 1)
	assert.Equal(t, "Toyota", cars[0].Brand)
}

func TestGetCarsByCity(t *testing.T) {
	m := new(mockRepo)
	expected := []*models.Car{{City: "Almaty"}}
	m.On("FindByCity", "Almaty").Return(expected, nil)

	uc := usecase.NewCarUseCase(m)
	cars, err := uc.GetCarsByCity("Almaty")

	assert.NoError(t, err)
	assert.Equal(t, "Almaty", cars[0].City)
}

func TestFindByModel(t *testing.T) {
	m := new(mockRepo)
	expected := []*models.Car{{Model: "Camry"}}
	m.On("FindByModel", "Camry").Return(expected, nil)

	uc := usecase.NewCarUseCase(m)
	cars, err := uc.FindByModel("Camry")

	assert.NoError(t, err)
	assert.Equal(t, "Camry", cars[0].Model)
}

func TestFindNearbyCars(t *testing.T) {
	m := new(mockRepo)
	expected := []*models.Car{{Latitude: 51.1, Longitude: 71.4}}
	m.On("FindNearby", 51.1, 71.4, 5.0).Return(expected, nil)

	uc := usecase.NewCarUseCase(m)
	cars, err := uc.FindNearbyCars(51.1, 71.4, 5.0)

	assert.NoError(t, err)
	assert.Equal(t, 51.1, cars[0].Latitude)
}

func TestChangeStatus(t *testing.T) {
	m := new(mockRepo)
	expected := &models.Car{Status: "occupied"}
	m.On("ChangeStatus", "abc123", "occupied").Return(expected, nil)

	uc := usecase.NewCarUseCase(m)
	car, err := uc.ChangeStatus(context.Background(), "abc123", "occupied")
	assert.NoError(t, err)
	assert.Equal(t, "occupied", car.Status)
}

func TestGetCarLocation(t *testing.T) {
	m := new(mockRepo)
	m.On("GetLocation", "carID").Return(51.1, 71.4, nil)

	uc := usecase.NewCarUseCase(m)
	lat, lon, err := uc.GetCarLocation("carID")

	assert.NoError(t, err)
	assert.Equal(t, 51.1, lat)
	assert.Equal(t, 71.4, lon)
}

func TestUpdateLocation(t *testing.T) {
	m := new(mockRepo)
	expected := &models.Car{Latitude: 50.0, Longitude: 70.0}
	m.On("UpdateLocation", "carID", 50.0, 70.0).Return(expected, nil)

	uc := usecase.NewCarUseCase(m)
	car, err := uc.UpdateLocation("carID", 50.0, 70.0)

	assert.NoError(t, err)
	assert.Equal(t, 50.0, car.Latitude)
	assert.Equal(t, 70.0, car.Longitude)
}
