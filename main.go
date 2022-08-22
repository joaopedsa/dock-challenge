package main

import (
	"log"

	"github.com/joaopedsa/dock-challenge/infra/database"
	"github.com/joaopedsa/dock-challenge/infra/server"
	"go.uber.org/zap"
)

func main() {
	if _, err := database.Connect(false); err != nil {
		log.Fatal(err)
	}
	if err := database.Ping(); err != nil {
		log.Fatal("Something is wrong with database database", zap.Error(err))
	}
	defer database.Close()
	server := server.NewServer()
	log.Fatalf("fatal error while running app server (%v)", server.Run())
}
