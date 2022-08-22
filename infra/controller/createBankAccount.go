package controller

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joaopedsa/dock-challenge/application"
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/utils"
)

type CreateBankAccountController struct {
	createBankAccount *application.CreateBankAccount
}

func NewCreateBankAccountController(createBankAccount *application.CreateBankAccount) *CreateBankAccountController {
	return &CreateBankAccountController{createBankAccount}
}

func (controller *CreateBankAccountController) Handle(c *fiber.Ctx) error {
	var bankAccount entity.BankAccount
	err := c.BodyParser(&bankAccount)
	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 400)
	}
	err = controller.createBankAccount.Execute(bankAccount)
	if strings.Contains(err.Error(), application.ErrBankAccountAlreadyExists.Error()) {
		log.Println(err)
		return utils.RespondWithError(c, err, 400)
	}
	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 500)
	}

	return c.Status(201).JSON("created bank account")
}
