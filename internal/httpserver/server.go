package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHandle() (http.Handler, error) {

	// новый роутер
	router := mux.NewRouter()

	//тут еще должен быть минимум коннект с БД

	// объявляем структуру с методами (путями)
	s := service{
		db: "connect DB",
	}

	// пути
	router.HandleFunc("/", s.home)
	router.HandleFunc("/about", s.about)

	return router, nil

}
