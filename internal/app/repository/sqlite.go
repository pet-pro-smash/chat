package repository

import "github.com/pet-pro-smash/chat/internal/app/httpserver/model"

type sqlite struct{}

func NewSqliteDB(c ConfigConnect) (*sqlite, error) {
	return &sqlite{}, nil
}

func (s *sqlite) CreateUser(user model.User) (int, error) {
	return 0, nil
}

func (s *sqlite) GetUser(email string) (model.User, error) {
	return model.User{}, nil
}

func (s *sqlite) DBClose() error {
	return nil
}
