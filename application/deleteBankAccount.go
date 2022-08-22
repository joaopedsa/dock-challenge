package application

import (
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/domain/repository"
)

type DeleteBankAccount struct {
	bankAccountRepository repository.BankAccountRepository
}

func NewDeleteBankAccount(bankAccountRepository repository.BankAccountRepository) *DeleteBankAccount {
	return &DeleteBankAccount{bankAccountRepository}
}

func (c *DeleteBankAccount) Execute(inputBankAccount entity.BankAccount) error {
	err := c.bankAccountRepository.Delete(inputBankAccount)
	if err != nil {
		return err
	}

	return nil
}
