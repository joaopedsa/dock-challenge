package controller

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joaopedsa/dock-challenge/application"
	"github.com/joaopedsa/dock-challenge/domain/entity"
	"github.com/joaopedsa/dock-challenge/utils"
)

type CreateUserController struct {
	createUser *application.CreateUser
}

func NewCreateUserController(createUser *application.CreateUser) *CreateUserController {
	return &CreateUserController{createUser}
}

func (controller *CreateUserController) Handle(c *fiber.Ctx) error {
	var user entity.User
	err := c.BodyParser(&user)
	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 400)
	}
	err = controller.createUser.Execute(user)
	if strings.Contains(err.Error(), application.ErrInvalidCPF.Error()) || strings.Contains(err.Error(), application.ErrUserAlreadyExists.Error()) {
		log.Println(err)
		return utils.RespondWithError(c, err, 400)
	}
	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 500)
	}

	return c.Status(201).JSON("created user")
}
