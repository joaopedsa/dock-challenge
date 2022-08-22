package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joaopedsa/dock-challenge/application"
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/utils"
)

type DeleteBankAccountController struct {
	deleteBankAccount *application.DeleteBankAccount
}

func NewDeleteBankAccountController(deleteBankAccount *application.DeleteBankAccount) *DeleteBankAccountController {
	return &DeleteBankAccountController{deleteBankAccount}
}

func (controller *DeleteBankAccountController) Handle(c *fiber.Ctx) error {
	var bankAccount entity.BankAccount
	err := c.QueryParser(&bankAccount)
	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 400)
	}

	if err := controller.deleteBankAccount.Execute(bankAccount); err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 500)
	}

	return c.Status(200).JSON("deleted bank account")
}
