package entity

import (
	"github.com/shopspring/decimal"
)

type BankStatement struct {
	BankAccountsID int32           `json:"bank_accounts_id"`
	Value          decimal.Decimal `json:"value"`
	Type           string          `json:"type"`
}

type ListBankStatement struct {
	Period         int32  `json:"period"`
	Agency         string `json:"agency"`
	Number         string `json:"number"`
	BankAccountsID int32  `json:"bank_accounts_id"`
	Type           string `json:"type"`
}
