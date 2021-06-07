package repository

type mysql struct{}

func NewMysqlDB(c ConfigConnect) (*mysql, error) {
	return &mysql{}, nil
}

func (ms *mysql) AddUser() {

}

func (ms *mysql) Close() error {
	return nil
}
