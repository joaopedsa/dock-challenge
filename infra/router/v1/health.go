package v1

import (
	"github.com/gofiber/fiber/v2"
)

func SetupV1Health(v1Router fiber.Router) error {
	// Health routes
	v1Router.Get("health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("running")
	})

	return nil
}
