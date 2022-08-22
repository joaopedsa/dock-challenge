package application

import (
	"strings"

	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/domain/repository"
	"github.com/shopspring/decimal"
)

type DepositBankAccount struct {
	bankAccountRepository   repository.BankAccountRepository
	bankStatementRepository repository.BankStatementRepository
}

func NewDepositBankAccount(bankAccountRepository repository.BankAccountRepository, bankStatementRepository repository.BankStatementRepository) *DepositBankAccount {
	return &DepositBankAccount{bankAccountRepository, bankStatementRepository}
}

func (w *DepositBankAccount) Execute(inputBankAccount entity.BankAccount, depositValue decimal.Decimal) error {
	bankAccount, err := w.bankAccountRepository.Get(inputBankAccount)
	if err != nil && strings.Contains(err.Error(), "record not found") {
		return ErrBankAccountNotFound
	}
	if err != nil {
		return err
	}
	if !bankAccount.IsActive {
		return ErrBankAccountIsNotActive
	}
	if bankAccount.IsBlocked {
		return ErrBankAccountBlocked
	}
	newBalance := bankAccount.Balance.Add(depositValue)
	if err := w.bankAccountRepository.UpdateBalance(inputBankAccount, newBalance); err != nil {
		return err
	}
	inputBankStatement := entity.BankStatement{
		BankAccountsID: bankAccount.ID,
		Value:          depositValue,
		Type:           "DEPOSIT",
	}
	if err := w.bankStatementRepository.Create(inputBankStatement); err != nil {
		w.bankAccountRepository.UpdateBalance(inputBankAccount, bankAccount.Balance)
		return err
	}

	return nil
}
