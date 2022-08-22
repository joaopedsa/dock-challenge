package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joaopedsa/dock-challenge/application"
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/utils"
)

type ListBankStatementController struct {
	listBankStatement *application.ListBankStatement
}

func NewListBankStatementController(listBankStatement *application.ListBankStatement) *ListBankStatementController {
	return &ListBankStatementController{listBankStatement}
}

func (controller *ListBankStatementController) Handle(c *fiber.Ctx) error {
	var inputBankStatement entity.ListBankStatement
	err := c.QueryParser(&inputBankStatement)
	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 400)
	}
	bankStatements, err := controller.listBankStatement.Execute(inputBankStatement)
	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 500)
	}

	return c.Status(200).JSON(bankStatements)
}
