package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pet-pro-smash/chat/internal/app/httpserver/model"
	"github.com/pet-pro-smash/chat/internal/app/httpserver/request"
)

// Обработчик регистрации
func (h *Handler) signUp(c *gin.Context) {
	var req request.SignUp
	if err := c.BindJSON(&req); err != nil {
		respondError(c, http.StatusBadRequest, errors.New("invalid input body"))
		return
	}

	if err := req.Validate(); err != nil {
		respondError(c, http.StatusBadRequest, err)
		return
	}

	user := model.User{Name: req.Name, Email: req.Email, Password: req.Password}
	id, err := h.service.Authorization.CreateUser(user)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err)
		return
	}

	respond(c, http.StatusCreated, gin.H{
		"id": id,
	})
}

// Обработчик аутентификации
func (h *Handler) signIn(c *gin.Context) {
	var req request.SignIn
	if err := c.BindJSON(&req); err != nil {
		respondError(c, http.StatusBadRequest, errors.New("invalid input body"))
		return
	}

	if err := req.Validate(); err != nil {
		respondError(c, http.StatusBadRequest, err)
		return
	}

	token, err := h.service.Authorization.GenerateToken(req.Email, req.Password)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err)
		return
	}

	respond(c, http.StatusOK, gin.H{
		"token": token,
	})
}
