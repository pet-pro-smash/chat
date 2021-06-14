package service

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pet-pro-smash/chat/internal/app/httpserver/model"
	"github.com/pet-pro-smash/chat/internal/app/repository"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenTTL   = time.Hour * 12
	signingKey = "SIGNING_KEY"
)

type AuthService struct {
	repos repository.Authorization
}

func NewAuthorizationService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func (s AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generateHashPassword(user.Password)

	return s.repos.CreateUser(user)
}

func (s AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repos.GetUser(email)
	if err != nil || !compareHashAndPassword(user.Password, password) {
		return "", errors.New("incorrect email or password")
	}

	type tokenClaims struct {
		jwt.StandardClaims
		UserId int `json:"user_id"`
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(os.Getenv(signingKey)))
}

func (s AuthService) ParseToken(accessToken string) (int, error) {
	type tokenClaims struct {
		jwt.StandardClaims
		UserId int `json:"user_id"`
	}

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, errors.New("invalid signing method")
		}

		return []byte(os.Getenv(signingKey)), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, err
	}

	return claims.UserId, nil
}

func generateHashPassword(password string) string {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(passwordHash)
}

func compareHashAndPassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
