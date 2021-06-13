package repository

import (
	"crypto/sha256"
	"fmt"
	"path/filepath"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sql struct {
	db *gorm.DB
}

type User struct {
	ID       int
	Name     string
	Email    string
	DateAdd  time.Time
	Password string
}

type Group struct {
	ID   int
	Name string
}

func hashPassword(p string) string {
	b := []byte(p)
	s := sha256.Sum256(b)
	return fmt.Sprintf("%x", s[:])
}

func NewSqliteDB(c ConfigConnect) (*sql, error) {
	// создание соединения с БД
	db, err := gorm.Open(sqlite.Open(filepath.Join("..", "data", "sqlite", "sql.db")), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к бд [sqlite]: %v", err)
	}

	// миграция
	if err := db.AutoMigrate(&User{}, &Group{}); err != nil {
		return nil, fmt.Errorf("ошибка миграции [sqlite]: %v", err)
	}

	r := &sql{db: db}

	if err := r.defaultCreate(); err != nil {
		return nil, fmt.Errorf(" error create default user")
	}

	return r, nil
}

func (ms *sql) defaultCreate() error {
	result := ms.db.Create(&User{ID: 1, Name: "Admin", Email: "admin@mail.com", DateAdd: time.Now(), Password: hashPassword("1234")})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ms *sql) CreateUser() {

}

func (ms *sql) DBClose() error {
	return nil
}
