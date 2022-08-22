package application

import (
	"github.com/joaopedsa/dock-challenge/domain/repository"
	"github.com/joaopedsa/dock-challenge/utils"
)

type DeleteUser struct {
	userRepository repository.UserRepository
}

func NewDeleteUser(userRepository repository.UserRepository) *DeleteUser {
	return &DeleteUser{userRepository}
}

type DeleteUserInput struct {
	CPF string
}

func (c *DeleteUser) Execute(input DeleteUserInput) error {
	input.CPF = utils.RemoveNonNumbersCPF(input.CPF)
	if !utils.ValidateCPF(input.CPF) {
		return ErrInvalidCPF
	}

	err := c.userRepository.Delete(input.CPF)
	if err != nil {
		return err
	}

	return nil
}
