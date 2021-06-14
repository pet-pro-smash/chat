package repository

import "github.com/pet-pro-smash/chat/internal/app/httpserver/model"

const (
	usersTable string = "users"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email string) (model.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db DBConnector) Repository {
	return Repository{Authorization: NewAuthSQL(db)}
}
