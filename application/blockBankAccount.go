package application

import (
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/domain/repository"
)

type BlockBankAccount struct {
	bankAccountRepository repository.BankAccountRepository
}

func NewBlockBankAccount(bankAccountRepository repository.BankAccountRepository) *BlockBankAccount {
	return &BlockBankAccount{bankAccountRepository}
}

func (b *BlockBankAccount) Execute(inputBankAccount entity.BankAccount) error {
	err := b.bankAccountRepository.UpdateBlock(inputBankAccount, true)
	if err != nil {
		return err
	}

	return nil
}
