package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
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

func (p *postgres) AddUser() {
	log.Println("postgres AddUser")
}

func (p *postgres) Close() error {
	if err := p.db.Close(); err != nil {
		return fmt.Errorf("произошла ошибка при закрытии соединения к бд: %s", err.Error())
	}
	return nil

}
