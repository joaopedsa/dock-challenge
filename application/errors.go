package application

import "errors"

var (
	ErrInvalidCPF               = errors.New("invalid CPF")
	ErrUserAlreadyExists        = errors.New("user already exists")
	ErrBankAccountNotFound      = errors.New("bank account not found")
	ErrInsuficientBalance       = errors.New("bank account insuficient balance")
	ErrBankAccountBlocked       = errors.New("bank account blocked")
	ErrBankAccountIsNotActive   = errors.New("bank account is not active")
	ErrBankAccountWithoutLimit  = errors.New("bank account without remaining limit")
	ErrBankAccountAlreadyExists = errors.New("bank account already exists")
)
