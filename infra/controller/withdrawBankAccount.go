package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joaopedsa/dock-challenge/application"
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/utils"
	"github.com/shopspring/decimal"
)

type WithdrawBankAccountController struct {
	WithdrawBankAccount *application.WithdrawBankAccount
}

func NewWithdrawBankAccountController(WithdrawBankAccount *application.WithdrawBankAccount) *WithdrawBankAccountController {
	return &WithdrawBankAccountController{WithdrawBankAccount}
}

func (controller *WithdrawBankAccountController) Handle(c *fiber.Ctx) error {
	var WithdrawBankAccount entity.WithdrawBankAccount
	err := c.BodyParser(&WithdrawBankAccount)
	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 400)
	}
	inputBankAccount := entity.BankAccount{
		Agency: WithdrawBankAccount.Agency,
		Number: WithdrawBankAccount.Number,
	}
	if WithdrawBankAccount.Value.LessThan(decimal.Zero) {
		log.Println(ErrInvalidParameter)
		return utils.RespondWithError(c, ErrInvalidParameter, 400)
	}
	if err := controller.WithdrawBankAccount.Execute(inputBankAccount, WithdrawBankAccount.Value); err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 500)
	}

	return c.Status(200).JSON("withdrew bank account")
}
