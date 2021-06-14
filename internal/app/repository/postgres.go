package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pet-pro-smash/chat/internal/app/httpserver/model"
)

type postgres struct {
	db *sqlx.DB
}

func NewPostgresDB(cfg ConfigConnect) (*postgres, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))

	return &postgres{db: db}, err
}

func (p *postgres) CreateUser(user model.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, email, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)

	err := p.db.QueryRow(query, user.Email, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		log.Println(err)
		return 0, errors.New("failed to create user")
	}

	return user.Id, nil
}

func (p *postgres) GetUser(email string) (model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT id, name, email, password_hash FROM %s ut WHERE ut.email = $1", usersTable)

	if err := p.db.QueryRow(query, email).Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
		return user, err
	}

	return user, nil
}

func (p *postgres) DBClose() error {
	if err := p.db.Close(); err != nil {
		return fmt.Errorf("произошла ошибка при закрытии соединения к бд: %s", err.Error())
	}
	return nil
}
