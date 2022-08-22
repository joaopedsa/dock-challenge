package application

import (
	"strings"

	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/domain/repository"
)

type GetBankAccount struct {
	bankAccountRepository repository.BankAccountRepository
}

func NewGetBankAccount(bankAccountRepository repository.BankAccountRepository) *GetBankAccount {
	return &GetBankAccount{bankAccountRepository}
}

func (g *GetBankAccount) Execute(inputBankAccount entity.BankAccount) (entity.BankAccount, error) {
	bankAccount, err := g.bankAccountRepository.Get(inputBankAccount)
	if err != nil && strings.Contains(err.Error(), "record not found") {
		return entity.BankAccount{}, ErrBankAccountNotFound
	}
	if err != nil {
		return entity.BankAccount{}, err
	}
	outputBankAccount := entity.BankAccount{
		Agency:    bankAccount.Agency,
		Number:    bankAccount.Number,
		Balance:   bankAccount.Balance,
		IsBlocked: bankAccount.IsBlocked,
	}

	return outputBankAccount, nil
}
