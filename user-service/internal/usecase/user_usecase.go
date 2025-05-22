package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/Akan15/carrental3/user-service/internal/models"
	natsPkg "github.com/Akan15/carrental3/user-service/internal/nats"
	"github.com/Akan15/carrental3/user-service/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo      repository.UserRepository
	sendEmail func(to, subject, body string) error
}

func NewUserUseCase(r repository.UserRepository, sendEmail func(to, subject, body string) error) *UserUseCase {
	return &UserUseCase{
		repo:      r,
		sendEmail: sendEmail,
	}
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
	message := fmt.Sprintf("Здравствуйте, %s!\nСпасибо за регистрацию в CarRental.", name)
	_ = u.sendEmail(email, "Добро пожаловать!", message)
	natsPkg.PublishUserCreated(email)

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

func (u *UserUseCase) ListUsers() ([]*models.User, error) {
	return u.repo.GetAllUsers()
}

func (u *UserUseCase) UpdateUser(id, name, email string) error {
	user, err := u.repo.GetUserByID(id)
	if err != nil {
		return err
	}

	user.Name = name
	user.Email = email

	return u.repo.UpdateUser(user)
}

func (u *UserUseCase) UpdatePassword(id, oldPassword, newPassword string) error {
	user, err := u.repo.GetUserByID(id)
	if err != nil {
		return err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)) != nil {
		return errors.New("old password does not match")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return u.repo.UpdateUserPassword(id, string(hashed))
}

func (u *UserUseCase) DeleteUser(id string) error {
	return u.repo.DeleteUser(id)
}

func (u *UserUseCase) UpdateEmail(id, password, newEmail string) error {
	user, err := u.repo.GetUserByID(id)
	if err != nil {
		return err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return errors.New("invalid password")
	}
	user.Email = newEmail
	return u.repo.UpdateUser(user)
}

func (u *UserUseCase) FindByName(name string) ([]*models.User, error) {
	return u.repo.FindByName(name)
}

func (u *UserUseCase) ChangePassword(email, oldPass, newPass string) error {
	user, err := u.repo.GetUserByEmail(email)
	if err != nil || user == nil {
		return errors.New("user not found")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPass)) != nil {
		return errors.New("old password incorrect")
	}
	newHash, _ := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	return u.repo.UpdateUserPassword(user.ID.Hex(), string(newHash))
}

func (u *UserUseCase) RefreshToken(oldToken string) (string, error) {
	userID, err := u.VerifyToken(oldToken)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	})
	return token.SignedString(jwtSecret)
}

func (u *UserUseCase) ResendVerification(email string) error {
	user, err := u.repo.GetUserByEmail(email)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	message := fmt.Sprintf("Здравствуйте, %s!\nПожалуйста, подтвердите свой email.", user.Name)
	return u.sendEmail(user.Email, "Подтверждение регистрации", message)
}
