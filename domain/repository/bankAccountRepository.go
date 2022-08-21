package repository

import "github.com/shopspring/decimal"

type BankAccountRepository interface {
	create(userID int32, number, agency string)
	delete(userID int32, number, agency string)
	block(userID int32, number, agency string, value decimal.Decimal)
	withdraw(userID int32, number, agency string, value decimal.Decimal)
	deposit(userID int32, number, agency string, value decimal.Decimal)
}
