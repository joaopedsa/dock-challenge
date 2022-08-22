package repository

import (
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/domain/models"
	"github.com/shopspring/decimal"
)

type BankStatementRepository interface {
	Create(inputBankStatement entity.BankStatement) error
	List(inputListBankStatement entity.ListBankStatement) ([]models.BankStatement, error)
	SumDeposits(inputListBankStatement entity.ListBankStatement) (decimal.Decimal, error)
}
