package usecase_test

import (
	"testing"

	"github.com/Akan15/carrental3/user-service/internal/models"
	"github.com/Akan15/carrental3/user-service/internal/usecase"
	"github.com/Akan15/carrental3/user-service/internal/usecase/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	// заглушка для отправки email
	dummySendEmail := func(to, subject, body string) error {
		return nil
	}

	uc := usecase.NewUserUseCase(mockRepo, dummySendEmail)

	user := &models.User{
		Name:     "John",
		Email:    "john@example.com",
		Password: "hashedpassword",
	}

	mockRepo.On("GetUserByEmail", user.Email).Return(nil, nil)
	mockRepo.On("CreateUser", mock.Anything).Return(nil)

	err := uc.Register(user.Name, user.Email, user.Password)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
