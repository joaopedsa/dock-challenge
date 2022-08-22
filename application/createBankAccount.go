package application

import (
	"strings"

	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/domain/repository"
)

type CreateBankAccount struct {
	bankAccountRepository repository.BankAccountRepository
}

func NewCreateBankAccount(bankAccountRepository repository.BankAccountRepository) *CreateBankAccount {
	return &CreateBankAccount{bankAccountRepository}
}

func (c *CreateBankAccount) Execute(inputBankAccount entity.BankAccount) error {
	err := c.bankAccountRepository.Create(inputBankAccount)
	if strings.Contains(err.Error(), "bank_accounts_number_agency_key") {
		return ErrBankAccountAlreadyExists
	}
	if err != nil {
		return err
	}

	return nil
}
