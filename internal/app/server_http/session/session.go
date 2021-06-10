package session

import "time"

// данные с сессиями
type sessionStore struct {
	store map[string]sesUser
}

// данные сессии одного клиента
type sesUser struct {
	dateCreate    time.Time // дата создания сессии
	dateLastVisit time.Time // время последнего использования
	sessionID     string    // индентификатор сессии UUID
	userID        string    // id user

}

// инициализируем сессии
func NewSeesion() *sessionStore {
	return &sessionStore{
		store: make(map[string]sesUser),
	}
}

// проверка наличия сессии по индификатору UUID
func (s *sessionStore) Check(string) (bool, error) {
	return true, nil
}

// создание новой сессии
func (s *sessionStore) Create() (string, error) {
	return "UUID", nil
}
