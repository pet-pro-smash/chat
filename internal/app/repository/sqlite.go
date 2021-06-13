package repository

type sqlite struct{}

func NewSqliteDB(c ConfigConnect) (*sqlite, error) {
	return &sqlite{}, nil
}

func (ms *sqlite) CreateUser() {

}

func (ms *sqlite) DBClose() error {
	return nil
}
