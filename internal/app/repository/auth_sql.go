package repository

import "github.com/pet-pro-smash/chat/internal/app/httpserver/model"

type AuthSQL struct {
	db DBConnector
}

func NewAuthSQL(db DBConnector) AuthSQL {
	return AuthSQL{db: db}
}

func (r AuthSQL) CreateUser(user model.User) (int, error) {
	return r.db.CreateUser(user)
}

func (r AuthSQL) GetUser(email string) (model.User, error) {
	return r.db.GetUser(email)
}
