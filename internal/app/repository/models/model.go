// модели для БД
package models

import "time"

type User struct {
	ID       int
	Name     string
	Email    string
	DateAdd  time.Time
	Password string
}

type Group struct {
	ID   int
	Name string
}
