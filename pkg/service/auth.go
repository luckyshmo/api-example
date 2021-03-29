package service

import (
	"github.com/luckyshmo/api-example/models"
	"github.com/luckyshmo/api-example/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	return 0, nil
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	return "token", nil
}

func (s *AuthService) ParseToken(token string) (int, error) {
	return 0, nil
}
