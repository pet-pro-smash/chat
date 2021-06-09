package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pet-pro-smash/chat/internal/app/server_http/service"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return Handler{service: service}
}

// middleware проверка куки и сессии
func session() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (h Handler) InitRoutes() *gin.Engine {
	// Созданеие роутера
	router := gin.New()

	router.Use()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pet-project")
	})

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	return router
}
