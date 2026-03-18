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
	db, err := InitDatabase()
	if err != nil {
		return err
	}
	repoLogin := repository.NewRepoLogin(db)
	repoUser := repository.NewRepoUser(db)

	// Use case
	login := usecase.NewLogin(repoLogin)
	user := usecase.NewUser(repoUser)

	// Handler
	handlerLogin := handler.NewLogin(login)
	handlerUser := handler.NewUser(user)

	// Router
	mux := http.NewServeMux()

	mux.HandleFunc("/login", handlerLogin.Login)
	mux.HandleFunc("/user/create", handlerUser.CreateUser)

	// Server
	return http.ListenAndServe(":8085", corsMiddleware(mux))
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

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		enableCORS(w)

		// Responde o preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
