package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type BankAccount struct {
	ID        int32           `json:"id"`
	UserID    int32           `json:"user_id"`
	Balance   decimal.Decimal `json:"balance"`
	Number    string          `json:"number"`
	Agency    string          `json:"agency"`
	IsActive  bool            `json:"is_active"`
	IsBlocked bool            `json:"is_blocked"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
