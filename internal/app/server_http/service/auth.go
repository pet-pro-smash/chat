package service

import "github.com/pet-pro-smash/chat/internal/app/repository"

type AuthorizationService struct {
	repos repository.Authorization
}

func NewAuthorizationService(repos repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repos: repos}
}

func (s AuthorizationService) CreateUser() {

}
