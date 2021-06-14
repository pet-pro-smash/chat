package repository

import "fmt"

// DBConnector | Интерфейс работы с базами данных, методы
type DBConnector interface {
	Authorization
	DBClose() error
}

// ConfigConnect | Конфиг соединения с БД
type ConfigConnect struct {
	Title    string
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// NewDBConnect | Инициализация соединения с БД
func NewDBConnect(cfg ConfigConnect) (DBConnector, error) {
	switch cfg.Title {
	case "postgres":
		p, err := NewPostgresDB(cfg)
		if err != nil {
			return nil, err
		}
		return p, nil

	case "sqlite":
		l, err := NewSqliteDB(cfg)
		if err != nil {
			return nil, err
		}
		return l, nil

	default:
		return nil, fmt.Errorf("база данных %s не найдена", cfg.Title)
	}
}
