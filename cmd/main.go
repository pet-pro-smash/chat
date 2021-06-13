package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/pet-pro-smash/chat/internal/app/cli"
	"github.com/pet-pro-smash/chat/internal/app/config"
	server_http "github.com/pet-pro-smash/chat/internal/app/httpserver"
	"github.com/pet-pro-smash/chat/internal/app/httpserver/handler"
	"github.com/pet-pro-smash/chat/internal/app/httpserver/service"
	"github.com/pet-pro-smash/chat/internal/app/repository"
	"github.com/sirupsen/logrus"
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
		logrus.Fatalf("ошибка при при подключении к бд: %v", err)
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
		signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
		<-done
		cancel()
	}()

	// Запуск http сервера
	wg.Add(1)
	go func() {
		defer wg.Done()

		// конфиг для сервера
		sc := server_http.Config{
			Host:           c.GetString("server_http.host"),
			Port:           c.GetString("server_http.port"),
			Handler:        handlers.InitRoutes(),
			MaxHeaderBytes: c.GetInt("server_http.max_header_bytes"),
		}

		// Инициализация http сервера
		srv := server_http.NewServer(sc)

		logrus.Infoln("запуск http сервера")

		// запуск сервера, блокирующая команда
		if err = srv.Start(ctx); err != nil {
			if err == http.ErrServerClosed {
				logrus.Println("HTTP сервер остановленн")
				return
			}
			logrus.Fatalf("произошла ошибка при запуске http сервера: %v", err)
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
	if err := db.DBClose(); err != nil {
		logrus.Error(err)
	}

	fmt.Println("Работа завершена!!!")
}
