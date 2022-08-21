package repository

type UserRepository interface {
	create(cpf, name string) error
	delete(cpf string) error
}
