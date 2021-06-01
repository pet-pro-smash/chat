package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ConfigPostgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg ConfigPostgres) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))

	return db, err
}
