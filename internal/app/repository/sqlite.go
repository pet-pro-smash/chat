package repository

type sqlite struct{}

func NewSqliteDB(c ConfigConnect) (*sqlite, error) {
	return &sqlite{}, nil
}

func (ms *sqlite) AddUser() {

}

func (ms *sqlite) Close() error {
	return nil
}
