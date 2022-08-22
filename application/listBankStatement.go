package application

import (
	"strings"

	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/domain/repository"
)

type ListBankStatement struct {
	bankStatementRepository repository.BankStatementRepository
	bankAccountRepository   repository.BankAccountRepository
}

func NewListBankStatement(bankStatementRepository repository.BankStatementRepository, bankAccountRepository repository.BankAccountRepository) *ListBankStatement {
	return &ListBankStatement{bankStatementRepository, bankAccountRepository}
}

func (l *ListBankStatement) Execute(inputListBankStatement entity.ListBankStatement) ([]entity.BankStatement, error) {
	inputBankAccount := entity.BankAccount{
		Number: inputListBankStatement.Number,
		Agency: inputListBankStatement.Agency,
	}
	bankAccount, err := l.bankAccountRepository.Get(inputBankAccount)
	if err != nil && strings.Contains(err.Error(), "record not found") {
		return []entity.BankStatement{}, ErrBankAccountNotFound
	}
	if err != nil {
		return []entity.BankStatement{}, err
	}
	inputListBankStatement.BankAccountsID = bankAccount.ID
	bankStatements, err := l.bankStatementRepository.List(inputListBankStatement)
	if err != nil {
		return []entity.BankStatement{}, err
	}

	var outputBankStatements []entity.BankStatement
	for _, bankStatement := range bankStatements {
		outputBankStatements = append(outputBankStatements, entity.BankStatement{
			Value: bankStatement.Value,
			Type:  bankStatement.Type,
		})
	}

	return outputBankStatements, nil
}
