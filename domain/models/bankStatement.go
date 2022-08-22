package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type BankStatement struct {
	ID            int32           `json:"id"`
	BankAccountID int32           `json:"bank_accounts_id"`
	Value         decimal.Decimal `json:"value"`
	Type          string          `json:"type"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

type SumBankStatement struct {
	TotalWithdraw decimal.Decimal `json:"total_withdraw"`
}
