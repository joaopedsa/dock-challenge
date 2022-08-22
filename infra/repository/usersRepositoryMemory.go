package repository

import (
	"errors"

	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/domain/models"
)

type UserRepositoryMemory struct {
	Users []models.User
}

func NewUserRepositoryMemory() *UserRepositoryMemory {
	return &UserRepositoryMemory{}
}

func (u *UserRepositoryMemory) Get(inputUser entity.User) (models.User, error) {
	for _, user := range u.Users {
		if user.CPF == inputUser.CPF {
			return user, nil
		}
	}

	return models.User{}, errors.New("record not found")
}

func (u *UserRepositoryMemory) Create(user entity.User) error {
	inputUser := models.User{
		CPF:  user.CPF,
		Name: user.Name,
	}
	u.Users = append(u.Users, inputUser)
	return nil
}

func (u *UserRepositoryMemory) Delete(cpf string) error {

	return nil
}
