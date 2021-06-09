package server_http

import (
	"context"
	"net/http"
	"time"
)

type Config struct {
	Host           string
	Port           string
	Handler        http.Handler
	MaxHeaderBytes int
}

type Server struct {
	httpServer *http.Server
}

// создаем сервер
func NewServer(cfg Config) Server {

	return Server{
		httpServer: &http.Server{
			Addr:           cfg.Host + ":" + cfg.Port,
			Handler:        cfg.Handler,
			MaxHeaderBytes: cfg.MaxHeaderBytes << 20,
			ReadTimeout:    time.Second * 10,
			WriteTimeout:   time.Second * 10,
		}}
}

// запуск сервера
func (s *Server) Start(ctx context.Context) error {

	// остановка сервера если появляется сигнал
	go func() {
		<-ctx.Done()
		s.httpServer.Shutdown(context.Background())
	}()
	return s.httpServer.ListenAndServe()
}

// мягкая остановка сервера
func (s *Server) shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
