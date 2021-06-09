package repository

import "fmt"

// интерфейс работы с базами данных, методы
type DBConnector interface {
	AddUser()
	Close() error
}

// конфиг соединения с БД
type ConfigConnect struct {
	Title    string
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// инициализация соединения с БД
func NewDBConnect(c ConfigConnect) (DBConnector, error) {
	switch c.Title {
	case "postgres":
		p, err := NewPostgresDB(c)
		if err != nil {
			return nil, err
		}
		return p, nil

	case "mysql":
		m, err := NewMysqlDB(c)
		if err != nil {
			return nil, err
		}
		return m, nil
	case "sqlite":
		l, err := NewMysqlDB(c)
		if err != nil {
			return nil, err
		}
		return l, nil
	default:
		return nil, fmt.Errorf(" Ошибка нет такой БД")
	}

}
