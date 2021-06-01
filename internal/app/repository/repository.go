package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
	CreateUser()
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{Authorization: NewAuthorizationSQL(db)}
}
