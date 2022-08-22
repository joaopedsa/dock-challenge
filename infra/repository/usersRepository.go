package repository

import (
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u UserRepository) Get(inputUser entity.User) (models.User, error) {
	var user models.User
	result := u.db.Table("users").Where("cpf = ?", inputUser.CPF).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (u UserRepository) Create(user entity.User) error {
	result := u.db.Table("users").Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u UserRepository) Delete(cpf string) error {
	result := u.db.Where("cpf = ?", cpf).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
