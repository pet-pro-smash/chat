package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pet-pro-smash/chat/internal/app/httpserver/service"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return Handler{service: service}
}

// InitRoutes | Инициализация роутера
func (h *Handler) InitRoutes() *gin.Engine {

	// режим запуска gin
	gin.SetMode(gin.ReleaseMode)
	// Созданеие роутера
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	return router
}
