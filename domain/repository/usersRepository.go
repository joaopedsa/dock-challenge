package repository

import (
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/domain/models"
)

type UserRepository interface {
	Get(inputUser entity.User) (models.User, error)
	Create(user entity.User) error
	Delete(cpf string) error
}
