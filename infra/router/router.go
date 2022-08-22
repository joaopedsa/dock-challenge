package router

import (
	"github.com/gofiber/fiber/v2"
	v1 "github.com/joaopedsa/dock-challenge/infra/router/v1"
)

type Router struct{}

func (r Router) SetupRouter(app *fiber.App) error {
	err := v1.SetupV1Router(app)
	if err != nil {
		return err
	}

	return nil
}
