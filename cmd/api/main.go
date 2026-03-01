package main

import (
	"net/http"
	"os"

	"github.com/edlanelima/system_sst/configs"
	"github.com/edlanelima/system_sst/internal/infraestructure/database"
	"github.com/edlanelima/system_sst/internal/infraestructure/repository"
	"github.com/edlanelima/system_sst/internal/interfaces/handler"
	"github.com/edlanelima/system_sst/internal/usecase"
)

func main() {

	if err := run(); err != nil {
		return
	}

}
func run() error {

	// Infra
	dbLogin, err := InitDatabase()
	if err != nil {
		return err
	}
	repoLogin := repository.NewRepoLogin(dbLogin)

	// Use case
	login := usecase.NewLogin(repoLogin)

	// Handler
	handlerLogin := handler.NewLogin(login)

	// Router
	http.HandleFunc("/login", handlerLogin.Login)

	// Server
	return http.ListenAndServe(":8080", nil)
}

func InitDatabase() (database.IDatabase, error) {
	scope := os.Getenv("scope")

	if scope == "" || scope == "local" {
		return database.NewMockDatabase(), nil
	}

	db, err := configs.NewMySQL()
	if err != nil {
		return nil, err
	}

	return database.NewDbMySQL(db), nil

}
