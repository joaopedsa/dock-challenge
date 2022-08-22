package utils

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorMessage struct {
	Error string `json:"message,omitempty"`
}

func RespondWithError(ctx *fiber.Ctx, err error, statusCode int) error {
	return ctx.Status(statusCode).JSON(ErrorMessage{
		Error: err.Error(),
	})
}
