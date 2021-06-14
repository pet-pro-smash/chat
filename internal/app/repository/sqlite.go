package repository

import (
	"crypto/sha256"
	"fmt"
	"path/filepath"
	"time"

	"github.com/pet-pro-smash/chat/internal/app/repository/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sql struct {
	db *gorm.DB
}

// просто тест, необращать внимание
func hashPassword(p string) string {
	b := []byte(p)
	s := sha256.Sum256(b)
	return fmt.Sprintf("%x", s[:])
}

// создаем соеинение с бд
func NewSqliteDB(c ConfigConnect) (*sql, error) {
	// создание соединения с БД
	db, err := gorm.Open(sqlite.Open(filepath.Join("..", "data", "sqlite", "sql.db")), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к бд [sqlite]: %v", err)
	}

	// миграция
	if err := db.AutoMigrate(&models.User{}, &models.Group{}); err != nil {
		return nil, fmt.Errorf("ошибка миграции [sqlite]: %v", err)
	}

	// соединение с БД
	r := &sql{db: db}

	// создание учетной записи администратора
	if err := r.defaultCreateAdmin(); err != nil {
		return nil, err
	}

	return r, nil
}

// создание администратора
// если БД создается при запуске, необходимо создать учетную запись администратора
func (ms *sql) defaultCreateAdmin() error {

	// смотрим учетную запись администратора
	result := ms.db.Where("name=?", "Admin").First(&models.User{})

	// если не найденна запись администратора, создаем новую
	if result.Error == gorm.ErrRecordNotFound {
		result := ms.db.Create(&models.User{ID: 1, Name: "Admin", Email: "admin@mail.com", DateAdd: time.Now(), Password: hashPassword("1234")})
		// если не удалось создать запись
		if result.Error != nil {
			return fmt.Errorf("ошибка создания записи администратора [sqlite]: %v", result.Error)
		}
		// запись успешно создан
		return nil
	}

	// проверяем на другие ошибки
	if result.Error != nil {
		return fmt.Errorf("ошибка проверки учетной записи администратора в БД [sqlite]: %v", result.Error)
	}

	// запись уже существует
	return nil
}

// создание нового пользователя
func (ms *sql) CreateUser() {

}

// завершение работы бд
func (ms *sql) DBClose() error {
	// получем куазател на БД от GORM
	sq, err := ms.db.DB()
	if err != nil {
		return fmt.Errorf("ошибка получения указателя sql.DB [sqlite]: %v", err)
	}

	// закрываем соединение
	if err := sq.Close(); err != nil {
		return fmt.Errorf("невозможно закрыть соединение с бд [sqlite]: %v", err)
	}

	return nil
}
