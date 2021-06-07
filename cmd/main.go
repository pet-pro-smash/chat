package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/pet-pro-smash/chat/internal/app/cli"
	"github.com/pet-pro-smash/chat/internal/app/config"
	"github.com/pet-pro-smash/chat/internal/app/handler"
	"github.com/pet-pro-smash/chat/internal/app/repository"
	"github.com/pet-pro-smash/chat/internal/app/server_http"
	"github.com/pet-pro-smash/chat/internal/app/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Завершение приложения при ошибке в инициализации конфигурационного файла
	c, err := config.LoadConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	// Инициализация бд postgres и завершение приложения при ошибке в подключении к бд
	db, err := repository.NewDBConnect(repository.ConfigConnect{
		Title:    c.GetString("db.title"),
		Host:     c.GetString("db.host"),
		Port:     c.GetString("db.port"),
		Username: c.GetString("db.username"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   c.GetString("db.db_name"),
		SSLMode:  c.GetString("db.ssl_mode"),
	})
	if err != nil {
		logrus.Fatalf("ошибка при при подключении к бд: %s", err.Error())
	}

	// Инициализация зависимостей
	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	// Ожидание на получение сигналов от системы для завершения работы
	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-done
		cancel()
	}()

	// Запуск http сервера
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Инициализация http сервера
		srv := server_http.NewServer(server_http.Config{
			Host:           viper.GetString("server.host"),
			Port:           viper.GetString("server.port"),
			Handler:        handlers.InitRoutes(),
			MaxHeaderBytes: viper.GetInt("server.max_header_bytes"),
		})

		logrus.Infoln("запуск http сервера")

		if err = srv.Start(ctx); err != nil {
			if err == http.ErrServerClosed {
				logrus.Println("HTTP сервер остановленн")
			}
			logrus.Fatalf("произошла ошибка при запуске http сервера: %s", err.Error())
		}

	}()

	// ws server
	wg.Add(1)
	go func() {
		defer wg.Done()
		// start server
	}()

	// tcp server
	wg.Add(1)
	go func() {
		defer wg.Done()
		// start server
	}()

	// простое консольное управление
	wg.Add(1)
	go func() {
		defer wg.Done()
		cli.CliStart(cli.Config{
			Ctx:    ctx,
			Cancel: cancel,
		})
	}()

	// ожидание завершения сервисов
	wg.Wait()

	// Закрытие соедниния с базой данных
	if err := db.Close(); err != nil {
		logrus.Error(err)
	}
}
