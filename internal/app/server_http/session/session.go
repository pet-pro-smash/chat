package session

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound = fmt.Errorf("ошибка [session]: такого ключа нет")
)

// хранилище с сессиями
type sessionStore struct {
	sync.Mutex
	store map[string]sesUser
}

// инициализируем сессии
func NewSeesion() *sessionStore {
	return &sessionStore{
		store: make(map[string]sesUser),
	}
}

// проверка наличия сессии по индификатору UUID
func (s *sessionStore) Get(g *gin.Context) (SessionWorker, error) {

	// получаем куку с идентификатором сессии
	c, err := g.Request.Cookie("s_id")
	if err != nil {
		return nil, fmt.Errorf("ошибка [session]: нет доступа к куки: %v", err)
	}

	s.Mutex.Lock()
	// проверяем есть ли сессия с таким индентификатором
	t, ok := s.store[c.Value]
	if !ok {
		return nil, ErrNotFound
	}

	// возвращаем данные о сессии
	return &t, nil
}

// создание новой сессии
func (s *sessionStore) Create() (string, error) {
	return "UUID", nil
}
