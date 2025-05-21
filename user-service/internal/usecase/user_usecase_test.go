package usecase_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/Akan15/carrental3/user-service/internal/models"
	"github.com/Akan15/carrental3/user-service/internal/usecase"
	"github.com/Akan15/carrental3/user-service/internal/usecase/mocks"
)

func TestRegister_UserAlreadyExists(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockRepo.On("GetUserByEmail", "test@example.com").Return(&models.User{Name: "Test"}, nil)

	uc := usecase.NewUserUseCase(mockRepo, nil)

	err := uc.Register("Test", "test@example.com", "1234")
	assert.EqualError(t, err, "user already exists")
}

func TestRegister_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockRepo.On("GetUserByEmail", "new@example.com").Return(nil, nil)
	mockRepo.On("CreateUser", mock.AnythingOfType("*models.User")).Return(nil)

	dummySendEmail := func(to, subject, body string) error {
		return nil
	}

	uc := usecase.NewUserUseCase(mockRepo, dummySendEmail)

	err := uc.Register("New", "new@example.com", "1234")
	assert.NoError(t, err)
}

func TestLogin_UserNotFound(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockRepo.On("GetUserByEmail", "missing@example.com").Return(nil, nil)

	uc := usecase.NewUserUseCase(mockRepo, nil)

	_, err := uc.Login("missing@example.com", "1234")
	assert.EqualError(t, err, "user not found")
}

func TestLogin_InvalidPassword(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockRepo.On("GetUserByEmail", "test@example.com").Return(&models.User{
		Email:    "test@example.com",
		Password: "$2a$10$invalidHashed",
	}, nil)

	uc := usecase.NewUserUseCase(mockRepo, nil)

	_, err := uc.Login("test@example.com", "wrongpass")
	assert.EqualError(t, err, "invalid credentials")
}

func TestUpdateUser_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	// Создаем валидный ObjectID и hex
	testObjID := primitive.NewObjectID()
	testHex := testObjID.Hex()

	mockRepo.On("GetUserByID", testHex).Return(&models.User{
		ID:    testObjID,
		Name:  "Old Name",
		Email: "old@example.com",
	}, nil)

	mockRepo.On("UpdateUser", mock.AnythingOfType("*models.User")).Return(nil)

	uc := usecase.NewUserUseCase(mockRepo, nil)

	err := uc.UpdateUser(testHex, "New Name", "new@example.com")
	assert.NoError(t, err)
}

func TestUpdatePassword_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	// Подготовим пользователя с валидным bcrypt-паролем: "oldpass"
	hashedOld, _ := bcrypt.GenerateFromPassword([]byte("oldpass"), bcrypt.DefaultCost)
	userID := primitive.NewObjectID()
	userHex := userID.Hex()

	mockRepo.On("GetUserByID", userHex).Return(&models.User{
		ID:       userID,
		Password: string(hashedOld),
	}, nil)

	// Проверим, что новый хеш передается в UpdateUserPassword
	mockRepo.On("UpdateUserPassword", userHex, mock.AnythingOfType("string")).Return(nil)

	uc := usecase.NewUserUseCase(mockRepo, nil)

	err := uc.UpdatePassword(userHex, "oldpass", "newpass")
	assert.NoError(t, err)
}

func TestDeleteUser_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	userID := primitive.NewObjectID().Hex()
	mockRepo.On("DeleteUser", userID).Return(nil)

	uc := usecase.NewUserUseCase(mockRepo, nil)

	err := uc.DeleteUser(userID)
	assert.NoError(t, err)
}

func TestUpdateEmail_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	userID := primitive.NewObjectID()
	userHex := userID.Hex()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("pass123"), bcrypt.DefaultCost)

	mockRepo.On("GetUserByID", userHex).Return(&models.User{
		ID:       userID,
		Password: string(hashedPassword),
	}, nil)

	mockRepo.On("UpdateUser", mock.AnythingOfType("*models.User")).Return(nil)

	uc := usecase.NewUserUseCase(mockRepo, nil)

	err := uc.UpdateEmail(userHex, "pass123", "new@email.com")
	assert.NoError(t, err)
}

func TestFindByName_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	mockRepo.On("FindByName", "Test").Return([]*models.User{
		{ID: primitive.NewObjectID(), Name: "Test", Email: "test@mail.com"},
	}, nil)

	uc := usecase.NewUserUseCase(mockRepo, nil)

	users, err := uc.FindByName("Test")
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}
func TestChangePassword_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	hashed, _ := bcrypt.GenerateFromPassword([]byte("oldpass"), bcrypt.DefaultCost)

	mockRepo.On("GetUserByEmail", "user@mail.com").Return(&models.User{
		ID:       primitive.NewObjectID(),
		Email:    "user@mail.com",
		Password: string(hashed),
	}, nil)

	mockRepo.On("UpdateUserPassword", mock.Anything, mock.Anything).Return(nil)

	uc := usecase.NewUserUseCase(mockRepo, nil)

	err := uc.ChangePassword("user@mail.com", "oldpass", "newpass")
	assert.NoError(t, err)
}
func TestRefreshToken_Success(t *testing.T) {
	// 🔒 Прямо создаем токен без Login()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": "testuserid123",
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	})

	tokenStr, err := token.SignedString([]byte("your_secret_key")) // 👈 тот же секрет, как в usecase
	assert.NoError(t, err)

	uc := usecase.NewUserUseCase(nil, nil)

	newToken, err := uc.RefreshToken(tokenStr)
	assert.NoError(t, err)
	assert.NotEmpty(t, newToken)
}
