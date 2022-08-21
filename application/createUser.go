package application

import "github.com/joaopedsa/dock-challenge/domain/repository"

type CreateUser struct {
	repository repository.UserRepository
}

func NewCreateUser(repository repository.UserRepository) *CreateUser {
	return &CreateUser{repository}
}

func (c *CreateUser) execute(input CreateUserInput) {
	c.repository.create(input.name, input.cpf)
}

type CreateUserInput struct {
	name string
	cpf  string
}
