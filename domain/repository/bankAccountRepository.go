package repository

import (
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/domain/models"
	"github.com/shopspring/decimal"
)

type BankAccountRepository interface {
	Get(inputBankAccount entity.BankAccount) (models.BankAccount, error)
	Create(inputBankAccount entity.BankAccount) error
	Delete(inputBankAccount entity.BankAccount) error
	UpdateBlock(inputBankAccount entity.BankAccount, newIsBlock bool) error
	UpdateBalance(inputBankAccount entity.BankAccount, newBalance decimal.Decimal) error
}
