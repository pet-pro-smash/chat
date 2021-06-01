package httpserver

import "net/http"

// структура для данных (соединение с БД и т.д.)
type service struct {
	db string
}

// GET /
func (s *service) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home ! /GET"))
}

// GET /about
func (s *service) about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Project  /GET"))
}
