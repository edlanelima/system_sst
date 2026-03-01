package usecase

import (
	"context"
	"errors"

	"github.com/edlanelima/system_sst/internal/domain/entity"
	"github.com/edlanelima/system_sst/internal/infraestructure/repository"
	"github.com/edlanelima/system_sst/internal/interfaces/dto"
)

type Login struct {
	repo repository.Login
}

func NewLogin(repo repository.Login) Login {
	return Login{repo: repo}
}

func (l Login) LoginLevel(ctx context.Context, params dto.Login) (entity.Role, error) {
	level, err := l.repo.SearchLevel(ctx, params.Email, params.Password)
	if err != nil {
		return "", err
	}

	if level == "" {
		return "", errors.New("user not found")
	}

	return level, nil
}
