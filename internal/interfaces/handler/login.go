package handler

import (
	"encoding/json"
	"net/http"

	"github.com/edlanelima/system_sst/internal/interfaces/dto"
	"github.com/edlanelima/system_sst/internal/usecase"
)

type LoginHandler struct {
	usecase usecase.Login
}

func NewLogin(usecase usecase.Login) LoginHandler {
	return LoginHandler{usecase}
}

func (l *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	login := dto.RequestLogin(r)

	if err := dto.IsValidParam(login); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	level, err := l.usecase.LoginLevel(ctx, login)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(level)
	return
}
