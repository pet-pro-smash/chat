package service

import (
	"github.com/pet-pro-smash/chat/internal/app/httpserver/model"
	"github.com/pet-pro-smash/chat/internal/app/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos repository.Repository) Service {
	return Service{Authorization: NewAuthorizationService(repos.Authorization)}
}
