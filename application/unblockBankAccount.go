package application

import (
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/domain/repository"
)

type UnblockBankAccount struct {
	bankAccountRepository repository.BankAccountRepository
}

func NewUnblockBankAccount(bankAccountRepository repository.BankAccountRepository) *UnblockBankAccount {
	return &UnblockBankAccount{bankAccountRepository}
}

func (u *UnblockBankAccount) Execute(inputBankAccount entity.BankAccount) error {
	err := u.bankAccountRepository.UpdateBlock(inputBankAccount, false)
	if err != nil {
		return err
	}

	return nil
}
