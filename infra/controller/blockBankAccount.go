package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joaopedsa/dock-challenge/application"
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/utils"
)

type BlockBankAccountController struct {
	blockBankAccount *application.BlockBankAccount
}

func NewBlockBankAccountController(blockBankAccount *application.BlockBankAccount) *BlockBankAccountController {
	return &BlockBankAccountController{blockBankAccount}
}

func (controller *BlockBankAccountController) Handle(c *fiber.Ctx) error {
	var bankAccount entity.BankAccount
	err := c.QueryParser(&bankAccount)
	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 400)
	}
	if err := controller.blockBankAccount.Execute(bankAccount); err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 500)
	}

	return c.Status(200).JSON("blocked bank account")
}
