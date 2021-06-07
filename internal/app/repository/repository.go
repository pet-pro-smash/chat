package repository

type Authorization interface {
	CreateUser()
}

type Repository struct {
	Authorization
}

func NewRepository(db DBConnector) Repository {
	return Repository{Authorization: NewAuthorizationSQL(db)}
}
