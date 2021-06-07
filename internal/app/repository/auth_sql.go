package repository

type AuthorizationSQL struct {
	db DBConnector
}

func NewAuthorizationSQL(db DBConnector) AuthorizationSQL {
	return AuthorizationSQL{db: db}
}

func (r AuthorizationSQL) CreateUser() {

}
