package repository

import (
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/domain/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BankAccountRepository struct {
	db *gorm.DB
}

func NewBankAccountRepository(db *gorm.DB) *BankAccountRepository {
	return &BankAccountRepository{
		db: db,
	}
}

func (b *BankAccountRepository) Get(inputBankAccount entity.BankAccount) (models.BankAccount, error) {
	var bankAccount models.BankAccount
	result := b.db.Table("bank_accounts").Where("number = ? and agency = ?", inputBankAccount.Number, inputBankAccount.Agency).First(&bankAccount)
	if result.Error != nil {
		return bankAccount, result.Error
	}

	return bankAccount, nil
}

func (b *BankAccountRepository) Create(inputBankAccount entity.BankAccount) error {
	result := b.db.Table("bank_accounts").Create(inputBankAccount)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BankAccountRepository) Delete(inputBankAccount entity.BankAccount) error {
	result := b.db.Table("bank_accounts").Where("agency = ? and number = ?", inputBankAccount.Agency, inputBankAccount.Number).Update("is_active", false)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BankAccountRepository) UpdateBlock(inputBankAccount entity.BankAccount, newIsBlocked bool) error {
	result := b.db.Table("bank_accounts").Where("agency = ? and number = ?", inputBankAccount.Agency, inputBankAccount.Number).Update("is_blocked", newIsBlocked)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BankAccountRepository) UpdateBalance(inputBankAccount entity.BankAccount, newBalance decimal.Decimal) error {
	result := b.db.Table("bank_accounts").Where("agency = ? and number = ?", inputBankAccount.Agency, inputBankAccount.Number).Update("balance", newBalance)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
