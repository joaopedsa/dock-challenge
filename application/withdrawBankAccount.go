package application

import (
	"strings"

	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/domain/repository"
	"github.com/shopspring/decimal"
)

type WithdrawBankAccount struct {
	bankAccountRepository   repository.BankAccountRepository
	bankStatementRepository repository.BankStatementRepository
}

func NewWithdrawBankAccount(bankAccountRepository repository.BankAccountRepository, bankStatementRepository repository.BankStatementRepository) *WithdrawBankAccount {
	return &WithdrawBankAccount{bankAccountRepository, bankStatementRepository}
}

func (w *WithdrawBankAccount) Execute(inputBankAccount entity.BankAccount, withdrawValue decimal.Decimal) error {
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

	inputListBankStatement := entity.ListBankStatement{
		Period:         1,
		BankAccountsID: bankAccount.ID,
		Type:           "WITHDRAW",
	}
	totalWithdrawExecuted, err := w.bankStatementRepository.SumDeposits(inputListBankStatement)
	if err != nil {
		return err
	}
	if totalWithdrawExecuted.Add(withdrawValue).GreaterThan(decimal.NewFromFloat(2000)) {
		return ErrBankAccountWithoutLimit
	}

	newBalance := bankAccount.Balance.Sub(withdrawValue)
	if newBalance.LessThan(decimal.Zero) {
		return ErrInsuficientBalance
	}
	if err := w.bankAccountRepository.UpdateBalance(inputBankAccount, newBalance); err != nil {
		return err
	}

	inputBankStatement := entity.BankStatement{
		BankAccountsID: bankAccount.ID,
		Value:          withdrawValue,
		Type:           "WITHDRAW",
	}
	if err := w.bankStatementRepository.Create(inputBankStatement); err != nil {
		w.bankAccountRepository.UpdateBalance(inputBankAccount, bankAccount.Balance)
		return err
	}

	return nil
}
