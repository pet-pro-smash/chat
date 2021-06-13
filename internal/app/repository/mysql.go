package repository

type mysql struct{}

func NewMysqlDB(c ConfigConnect) (*mysql, error) {
	return &mysql{}, nil
}

func (ms *mysql) CreateUser() {

}

func (ms *mysql) DBClose() error {
	return nil
}
