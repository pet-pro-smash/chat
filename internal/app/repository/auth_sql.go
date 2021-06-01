package repository

import "github.com/jmoiron/sqlx"

type AuthorizationSQL struct {
	db *sqlx.DB
}

func NewAuthorizationSQL(db *sqlx.DB) AuthorizationSQL {
	return AuthorizationSQL{db: db}
}

func (r AuthorizationSQL) CreateUser() {

}
