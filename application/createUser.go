package application

import (
	"strings"

	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/domain/repository"
	"github.com/joaopedsa/dock-challenge/utils"
)

type CreateUser struct {
	userRepository repository.UserRepository
}

func NewCreateUser(userRepository repository.UserRepository) *CreateUser {
	return &CreateUser{userRepository}
}

func (c *CreateUser) Execute(inputUser entity.User) error {
	inputUser.CPF = utils.RemoveNonNumbersCPF(inputUser.CPF)
	if !utils.ValidateCPF(inputUser.CPF) {
		return ErrInvalidCPF
	}
	user, err := c.userRepository.Get(inputUser)
	if user.CPF != "" {
		return ErrUserAlreadyExists
	}
	if err != nil && !strings.Contains(err.Error(), "record not found") {
		return err
	}
	err = c.userRepository.Create(inputUser)
	if err != nil {
		return err
	}

	return nil
}
