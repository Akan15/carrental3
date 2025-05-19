package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/Akan15/carrental3/user-service/internal/models"
	"github.com/Akan15/carrental3/user-service/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

func (u *UserUseCase) Register(name, email, password string) error {
	existingUser, _ := u.repo.GetUserByEmail(email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := u.repo.CreateUser(user); err != nil {
		return err
	}

	// ⬇️ Отправка письма
	message := fmt.Sprintf("Здравствуйте, %s!\nСпасибо за регистрацию в CarRental Akan.", name)
	_ = SendEmail(email, "Добро пожаловать!", message)

	return nil
}

var jwtSecret = []byte("your_secret_key") // для безопасности вынеси в .env

func (u *UserUseCase) Login(email, password string) (string, error) {
	user, err := u.repo.GetUserByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.Hex(),
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (u *UserUseCase) VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Убедимся, что алгоритм правильный
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["user_id"] == nil {
		return "", errors.New("invalid claims")
	}

	return claims["user_id"].(string), nil
}
