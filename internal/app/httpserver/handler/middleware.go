package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	userCtx string = "userId"
)

func (h *Handler) userAuthentication(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		respondError(c, http.StatusUnauthorized, errors.New("empty auth header"))
		return
	}

	headerValue := strings.Split(header, " ")
	if len(headerValue) != 2 {
		respondError(c, http.StatusUnauthorized, errors.New("invalid auth header"))
		return
	}

	if headerValue[0] != "Bearer" {
		respondError(c, http.StatusUnauthorized, errors.New("invalid auth header"))
		return
	}

	if headerValue[1] == "" {
		respondError(c, http.StatusUnauthorized, errors.New("token is empty"))
		return
	}

	id, err := h.service.Authorization.ParseToken(headerValue[1])
	if err != nil {
		respondError(c, http.StatusUnauthorized, err)
		return
	}

	c.Set(userCtx, id)
}

//// middleware проверка куки и сессии
//func (h *Handler) session() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// проверка куки
//		// проверка сессии
//		fmt.Println("MIDDLEWARE")
//		c.Next()
//	}
//}
