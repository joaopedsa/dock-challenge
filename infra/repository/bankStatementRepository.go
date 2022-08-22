package repository

import (
	"time"

	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/domain/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BankStatementRepository struct {
	db *gorm.DB
}

func NewBankStatementRepository(db *gorm.DB) *BankStatementRepository {
	return &BankStatementRepository{
		db: db,
	}
}

func (b *BankStatementRepository) Create(inputBankStatement entity.BankStatement) error {
	result := b.db.Table("bank_statements").Create(inputBankStatement)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BankStatementRepository) List(inputBankStatement entity.ListBankStatement) ([]models.BankStatement, error) {
	var bankStatements []models.BankStatement
	initialDate := time.Now().AddDate(0, 0, -int(inputBankStatement.Period))

	result := b.db.Table("bank_statements").Where("bank_accounts_id = ? and created_at > ?", inputBankStatement.BankAccountsID, initialDate).Find(&bankStatements)
	if result.Error != nil {
		return nil, result.Error
	}

	return bankStatements, nil
}

func (b *BankStatementRepository) SumDeposits(inputBankStatement entity.ListBankStatement) (decimal.Decimal, error) {
	initialDate := time.Now().AddDate(0, 0, -int(inputBankStatement.Period))
	var sumBankStatement models.SumBankStatement
	result := b.db.Table("bank_statements").Select("sum(value) as total_withdraw").Where("bank_accounts_id = ? and created_at > ? and type = ?", inputBankStatement.BankAccountsID, initialDate, "WITHDRAW").Scan(&sumBankStatement)
	if result.Error != nil {
		return sumBankStatement.TotalWithdraw, result.Error
	}

	return sumBankStatement.TotalWithdraw, nil
}
