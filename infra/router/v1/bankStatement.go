package v1

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joaopedsa/dock-challenge/application"
	"github.com/joaopedsa/dock-challenge/infra/controller"
	"github.com/joaopedsa/dock-challenge/infra/database"
	"github.com/joaopedsa/dock-challenge/infra/repository"
)

type bankStatementControllers struct {
	listBankStatementController *controller.ListBankStatementController
}

func SetupBankStatementControllers() bankStatementControllers {
	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	bankStatementRepository := repository.NewBankStatementRepository(db)
	bankAccountRepository := repository.NewBankAccountRepository(db)

	listBankStatement := application.NewListBankStatement(bankStatementRepository, bankAccountRepository)
	listBankStatementController := controller.NewListBankStatementController(listBankStatement)

	return bankStatementControllers{
		listBankStatementController: listBankStatementController,
	}
}

func SetupV1BankStatement(v1Router fiber.Router) error {
	bankstatementControllers := SetupBankStatementControllers()

	// bank statement routes
	v1Router.Get("bank-statement", bankstatementControllers.listBankStatementController.Handle)

	return nil
}
