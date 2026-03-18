package dto

import (
	"errors"
	"net/http"
)

type Login struct {
	Email    string
	Password string
}

func RequestLogin(r *http.Request) Login {
	return Login{
		Email:    r.URL.Query().Get("email"),
		Password: r.URL.Query().Get("password"),
	}
}

func IsValidParam(input Login) error {
	if input.Email == "" || input.Password == "" {
		return errors.New("email and login are mandatory parameters")
	}
	return nil
}
