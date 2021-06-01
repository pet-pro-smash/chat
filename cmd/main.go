package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pet-pro-smash/chat/internal/app/handler"
	"github.com/pet-pro-smash/chat/internal/app/repository"
	"github.com/pet-pro-smash/chat/internal/app/server"
	"github.com/pet-pro-smash/chat/internal/app/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// Инициализация конфигурационного файла
func initConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	// Завершение приложения при ошибке в инициализации конфигурационного файла
	if err := initConfig(); err != nil {
		logrus.Fatalf("ошибка при инициализации конфигурационного файла: %s", err.Error())
	}

	// Завершение приложения при ошибке в загрузке переменных окружения
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("ошибка при загрузке переменных окружения: %s", err.Error())
	}

	// Инициализация бд postgres и завершение приложения при ошибке в подключении к бд
	db, err := repository.NewPostgresDB(repository.ConfigPostgres{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   viper.GetString("db.db_name"),
		SSLMode:  viper.GetString("db.ssl_mode"),
	})
	if err != nil {
		logrus.Fatalf("ошибка при при подключении к бд: %s", err.Error())
	}

	// Инициализация зависимостей
	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	// Инициализация http сервера
	srv := server.NewServer(server.Config{
		Host:           viper.GetString("server.host"),
		Port:           viper.GetString("server.port"),
		Handler:        handlers.InitRoutes(),
		MaxHeaderBytes: viper.GetInt("server.max_header_bytes"),
	})

	// Запуск сервера в go-рутине (для плавной остановки сервера)
	go func() {
		if err = srv.Start(); err != nil {
			logrus.Fatalf("произошла ошибка при запуске http сервера: %s", err.Error())
		}
	}()

	logrus.Infoln("запуск http сервера")

	// Чтение из канала двух системных сигналов
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("завершение работы http сервера")

	// Закрытие соедниния с базой данных
	if err := db.Close(); err != nil {
		logrus.Errorf("произошла ошибка при закрытии соединения к бд: %s", err.Error())
	}

	// Плавная остановка сервера
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("произошла ошибка при остановке сервера: %s", err.Error())
	}
}
