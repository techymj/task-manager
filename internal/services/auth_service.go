package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/techymj/task-manager/internal/models"
	"github.com/techymj/task-manager/internal/repositories"
)

type AuthService struct {
	UserRepo repositories.UserRepository
	JWTKey   []byte
}

func NewAuthService(repo repositories.UserRepository, key string) *AuthService {
	return &AuthService{
		UserRepo: repo,
		JWTKey:   []byte(key),
	}
}

func (s *AuthService) Register(user *models.User) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)
	return s.UserRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.UserRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.JWTKey)
}
