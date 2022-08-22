package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joaopedsa/dock-challenge/application"
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/utils"
	"github.com/shopspring/decimal"
)

type DepositBankAccountController struct {
	DepositBankAccount *application.DepositBankAccount
}

func NewDepositBankAccountController(DepositBankAccount *application.DepositBankAccount) *DepositBankAccountController {
	return &DepositBankAccountController{DepositBankAccount}
}

func (controller *DepositBankAccountController) Handle(c *fiber.Ctx) error {
	var depositBankAccount entity.DepositBankAccount
	err := c.BodyParser(&depositBankAccount)
	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 400)
	}
	inputBankAccount := entity.BankAccount{
		Agency: depositBankAccount.Agency,
		Number: depositBankAccount.Number,
	}
	if depositBankAccount.Value.LessThan(decimal.Zero) {
		log.Println(ErrInvalidParameter)
		return utils.RespondWithError(c, ErrInvalidParameter, 400)
	}
	if err := controller.DepositBankAccount.Execute(inputBankAccount, depositBankAccount.Value); err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 500)
	}

	return c.Status(200).JSON("deposited bank account")
}
