package session

import (
	"time"

	"github.com/gin-gonic/gin"
)

type SessionWorker interface {
	Save() error
	Delete() error
}

// данные сессии одного клиента
type sesUser struct {
	dateCreate    time.Time    // дата создания сессии
	dateLastVisit time.Time    // время последнего использования
	sessionID     string       // индентификатор сессии UUID
	userID        string       // id user
	gin           *gin.Context // данные о сеансе

}

func (s *sesUser) Save() error {
	return nil
}

func (s *sesUser) Delete() error {
	return nil
}
