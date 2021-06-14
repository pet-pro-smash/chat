package httpserver

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

// Запуск сервера
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// Плавная остановка сервера
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
