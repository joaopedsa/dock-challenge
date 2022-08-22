package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joaopedsa/dock-challenge/application"
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/utils"
)

type DeleteUserController struct {
	deleteUser *application.DeleteUser
}

func NewDeleteUserController(deleteUserApplication *application.DeleteUser) *DeleteUserController {
	return &DeleteUserController{deleteUserApplication}
}

func (controller *DeleteUserController) Handle(c *fiber.Ctx) error {
	var user entity.User
	if err := c.QueryParser(&user); err != nil {
		log.Println(err)
		return utils.RespondWithError(c, ErrInvalidParameter, 400)
	}
	inputDeleteUser := application.DeleteUserInput{
		CPF: user.CPF,
	}
	if err := controller.deleteUser.Execute(inputDeleteUser); err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 500)
	}

	return c.Status(200).JSON("deleted user")
}
