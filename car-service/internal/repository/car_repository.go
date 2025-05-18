package repository

import "github.com/Akan15/carrental3/car-service/internal/models"

type CarRepository interface {
	GetCarByID(id string) (*models.Car, error)
}

type InMemoryCarRepo struct {
	data map[string]*models.Car
}

func NewInMemoryCarRepo() *InMemoryCarRepo {
	return &InMemoryCarRepo{
		data: map[string]*models.Car{
			"1": {ID: "1", Brand: "Toyota", Model: "Camry"},
			"2": {ID: "2", Brand: "Tesla", Model: "Model 3"},
		},
	}
}

func (r *InMemoryCarRepo) GetCarByID(id string) (*models.Car, error) {
	car, exists := r.data[id]
	if !exists {
		return nil, nil
	}
	return car, nil
}
