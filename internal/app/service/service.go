package service

import "github.com/pet-pro-smash/chat/internal/app/repository"

type Authorization interface {
	CreateUser()
}

type Service struct {
	Authorization
}

func NewService(repos repository.Repository) Service {
	return Service{Authorization: repos.Authorization}
}
