package application

import (
	"testing"

	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	userRepository := repository.NewUserRepositoryMemory()
	createUser := NewCreateUser(userRepository)
	expectedInvalidCPF := entity.User{
		Name: "Test",
		CPF:  "INVALID CPF",
	}
	assert.Equal(ErrInvalidCPF.Error(), createUser.Execute(expectedInvalidCPF).Error())

	expectedSuccess := entity.User{
		Name: "João Pedro Santana",
		CPF:  "09662672907",
	}
	assert.Nil(createUser.Execute(expectedSuccess))

	expectedAlreadyExists := entity.User{
		Name: "João Pedro Santana",
		CPF:  "09662672907",
	}
	assert.Equal(ErrUserAlreadyExists.Error(), createUser.Execute(expectedAlreadyExists).Error())
}
