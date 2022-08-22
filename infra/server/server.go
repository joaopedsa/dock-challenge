package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joaopedsa/dock-challenge/infra/router"
	"github.com/joaopedsa/dock-challenge/utils"
)

type Server struct {
	App    *fiber.App
	Router router.Router
}

func NewServer() Server {
	app := fiber.New()

	return Server{
		App:    app,
		Router: router.Router{},
	}
}

func (s Server) Run() error {
	err := s.Router.SetupRouter(s.App)
	if err != nil {
		return err
	}

	port := utils.GetEnv("SERVER_PORT", "8080")
	return s.App.Listen(fmt.Sprintf(":%v", port))
}
