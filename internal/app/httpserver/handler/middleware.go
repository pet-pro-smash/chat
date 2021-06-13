package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// middleware проверка куки и сессии
func (h *Handler) session() gin.HandlerFunc {
	return func(c *gin.Context) {
		// проверка куки
		// проверка сессии
		fmt.Println("MIDDLEWARE")
		c.Next()
	}
}
