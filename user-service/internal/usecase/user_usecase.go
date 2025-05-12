package usecase

import (
	"user-service/internal/models"
	"user-service/internal/repository"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (u *UserUseCase) GetUserByID(id string) (*models.User, error) {
	return u.repo.GetUserByID(id)
}
