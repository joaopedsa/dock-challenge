package controller

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joaopedsa/dock-challenge/application"
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/utils"
)

type GetBankAccountController struct {
	getBankAccount *application.GetBankAccount
}

func NewGetBankAccountController(getBankAccount *application.GetBankAccount) *GetBankAccountController {
	return &GetBankAccountController{getBankAccount}
}

func (controller *GetBankAccountController) Handle(c *fiber.Ctx) error {
	var bankAccount entity.BankAccount
	err := c.QueryParser(&bankAccount)
	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 400)
	}
	bankAccount, err = controller.getBankAccount.Execute(bankAccount)
	if err != nil && strings.Contains(err.Error(), application.ErrBankAccountNotFound.Error()) {
		log.Println(err)
		return utils.RespondWithError(c, err, 404)
	}

	return c.Status(200).JSON(bankAccount)
}
