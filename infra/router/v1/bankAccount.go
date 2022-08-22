package v1

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joaopedsa/dock-challenge/application"
	"github.com/joaopedsa/dock-challenge/infra/controller"
	"github.com/joaopedsa/dock-challenge/infra/database"
	"github.com/joaopedsa/dock-challenge/infra/repository"
)

type bankAccountControllers struct {
	createBankAccount   *controller.CreateBankAccountController
	deleteBankAccount   *controller.DeleteBankAccountController
	getBankAccount      *controller.GetBankAccountController
	blockBankAccount    *controller.BlockBankAccountController
	unblockBankAccount  *controller.UnblockBankAccountController
	depositBankAccount  *controller.DepositBankAccountController
	withdrawBankAccount *controller.WithdrawBankAccountController
}

func SetupV1BankAccount(v1Router fiber.Router) error {
	bankAccountControllers := SetupBankAccountControllers()

	// bankAccount routes
	v1Router.Post("bank-account", bankAccountControllers.createBankAccount.Handle)
	v1Router.Delete("bank-account", bankAccountControllers.deleteBankAccount.Handle)
	v1Router.Get("bank-account", bankAccountControllers.getBankAccount.Handle)
	v1Router.Get("bank-account/block", bankAccountControllers.blockBankAccount.Handle)
	v1Router.Get("bank-account/unblock", bankAccountControllers.unblockBankAccount.Handle)
	v1Router.Post("bank-account/deposit", bankAccountControllers.depositBankAccount.Handle)
	v1Router.Post("bank-account/withdraw", bankAccountControllers.withdrawBankAccount.Handle)

	return nil
}

func SetupBankAccountControllers() bankAccountControllers {
	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	bankAccountRepository := repository.NewBankAccountRepository(db)
	bankStatementRepository := repository.NewBankStatementRepository(db)

	createBankAccount := application.NewCreateBankAccount(bankAccountRepository)
	createBankAccountController := controller.NewCreateBankAccountController(createBankAccount)
	deleteBankAccount := application.NewDeleteBankAccount(bankAccountRepository)
	deleteBankAccountController := controller.NewDeleteBankAccountController(deleteBankAccount)
	getBankAccount := application.NewGetBankAccount(bankAccountRepository)
	getBankAccountController := controller.NewGetBankAccountController(getBankAccount)
	blockBankAccount := application.NewBlockBankAccount(bankAccountRepository)
	blockBankAccountController := controller.NewBlockBankAccountController(blockBankAccount)
	unblockBankAccount := application.NewUnblockBankAccount(bankAccountRepository)
	unblockBankAccountController := controller.NewUnblockBankAccountController(unblockBankAccount)
	depositBankAccount := application.NewDepositBankAccount(bankAccountRepository, bankStatementRepository)
	DepositBankAccountController := controller.NewDepositBankAccountController(depositBankAccount)
	withdrawBankAccount := application.NewWithdrawBankAccount(bankAccountRepository, bankStatementRepository)
	WithdrawBankAccountController := controller.NewWithdrawBankAccountController(withdrawBankAccount)

	return bankAccountControllers{
		createBankAccount:   createBankAccountController,
		deleteBankAccount:   deleteBankAccountController,
		getBankAccount:      getBankAccountController,
		blockBankAccount:    blockBankAccountController,
		unblockBankAccount:  unblockBankAccountController,
		depositBankAccount:  DepositBankAccountController,
		withdrawBankAccount: WithdrawBankAccountController,
	}
}
