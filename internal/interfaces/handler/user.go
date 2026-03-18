package handler

import (
	"encoding/json"
	"net/http"

	"github.com/edlanelima/system_sst/internal/interfaces/dto"
	"github.com/edlanelima/system_sst/internal/usecase"
)

type UserHandler struct {
	usecase usecase.User
}

func NewUser(usecase usecase.User) UserHandler {
	return UserHandler{usecase}
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user dto.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": err.Error(),
			"user":    user,
		})
		return
	}

	if err := u.usecase.CreateUser(ctx, user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": err.Error(),
			"user":    user,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Usuário criado com sucesso",
		"user":    user,
	})
}
