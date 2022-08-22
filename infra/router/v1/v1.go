package v1

import "github.com/gofiber/fiber/v2"

func SetupV1Router(app *fiber.App) error {
	v1Routes := app.Group("/v1")
	SetupV1Health(v1Routes)
	SetupV1Users(v1Routes)
	SetupV1BankAccount(v1Routes)
	SetupV1BankStatement(v1Routes)

	return nil
}
