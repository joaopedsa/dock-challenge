package repository

import "github.com/shopspring/decimal"

type BankStatementRepository interface {
	create(bankAccountID int32, typeStatement string, value decimal.Decimal) error
	list(bankAccountID int32, periodTime string) ([]BankStatement, error)
}
