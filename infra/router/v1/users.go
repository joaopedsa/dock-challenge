package v1

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joaopedsa/dock-challenge/application"
	"github.com/joaopedsa/dock-challenge/infra/controller"
	"github.com/joaopedsa/dock-challenge/infra/database"
	"github.com/joaopedsa/dock-challenge/infra/repository"
)

type userControllers struct {
	createUserController *controller.CreateUserController
	deleteUserController *controller.DeleteUserController
}

func SetupUserControllers() userControllers {
	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	userRepository := repository.NewUserRepository(db)
	createUserApplication := application.NewCreateUser(userRepository)
	createUserController := controller.NewCreateUserController(createUserApplication)
	deleteUserApplication := application.NewDeleteUser(userRepository)
	deleteUserController := controller.NewDeleteUserController(deleteUserApplication)

	return userControllers{
		createUserController: createUserController,
		deleteUserController: deleteUserController,
	}
}

func SetupV1Users(v1Router fiber.Router) error {
	userControllers := SetupUserControllers()

	// users routes
	v1Router.Post("users", userControllers.createUserController.Handle)
	v1Router.Delete("users", userControllers.deleteUserController.Handle)

	return nil
}
