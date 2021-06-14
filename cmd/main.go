package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	httpSrv "github.com/pet-pro-smash/chat/internal/app/httpserver"
	"github.com/pet-pro-smash/chat/internal/app/httpserver/handler"
	"github.com/pet-pro-smash/chat/internal/app/httpserver/service"
	"github.com/pet-pro-smash/chat/internal/app/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Завершение приложения при ошибке в инициализации конфигурационного файла
	if err := initConfig(); err != nil {
		logrus.Fatalf("произошла ошибка при иниализации конфигурационного файла: %s", err.Error())
	}

	// Инициализация бд postgres и завершение приложения при ошибке в подключении к бд
	db, err := repository.NewDBConnect(repository.ConfigConnect{
		Title:    viper.GetString("db.title"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   viper.GetString("db.db_name"),
		SSLMode:  viper.GetString("db.ssl_mode"),
	})
	if err != nil {
		logrus.Fatalf("ошибка при при подключении к бд: %v", err)
	}

	// Инициализация зависимостей
	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	// Запуск http сервера
	// конфиг для сервера
	srvConfig := httpSrv.Config{
		Host:           viper.GetString("server_http.host"),
		Port:           viper.GetString("server_http.port"),
		Handler:        handlers.InitRoutes(),
		MaxHeaderBytes: viper.GetInt("server_http.max_header_bytes"),
	}

	// Инициализация http сервера
	srv := httpSrv.NewServer(srvConfig)

	logrus.Infoln("запуск http сервера")

	// запуск сервера, блокирующая команда
	if err = srv.Start(); err != nil {
		logrus.Fatalf("произошла ошибка при запуске http сервера: %v", err)
	}

	// Закрытие соедниния с базой данных
	if err := db.DBClose(); err != nil {
		logrus.Error(err)
	}

	fmt.Println("Приложение остановлено")
}

func initConfig() error {
	viper.AddConfigPath("config/")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}
