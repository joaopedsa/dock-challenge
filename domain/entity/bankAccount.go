package entity

import "github.com/shopspring/decimal"

type BankAccount struct {
	UserID    int32           `json:"user_id"`
	Balance   decimal.Decimal `json:"balance"`
	Number    string          `json:"number"`
	Agency    string          `json:"agency"`
	IsBlocked bool            `json:"is_blocked"`
}

type DepositBankAccount struct {
	Value  decimal.Decimal `json:"value"`
	Number string          `json:"number"`
	Agency string          `json:"agency"`
}

type WithdrawBankAccount struct {
	Value  decimal.Decimal `json:"value"`
	Number string          `json:"number"`
	Agency string          `json:"agency"`
}
