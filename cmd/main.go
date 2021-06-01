package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pet-pro-smash/chat/internal/config"
	"github.com/pet-pro-smash/chat/internal/httpserver"
)

func main() {

	if err := start(); err != nil {
		log.Fatal(err)
	}

	log.Println("Работа программы завершена")

}

func start() error {

	// загружаем конфиг
	conf, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("Не удалось получить конфиг")
	}

	// получаем Handler
	h, err := httpserver.NewHandle()
	if err != nil {
		return err
	}

	// конфигурируем сервер
	s := http.Server{
		Addr:    fmt.Sprintf(":%d", conf.HttpServer.Port),
		Handler: h,
	}

	// запускаем сервер
	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
