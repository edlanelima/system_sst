package usecase

import (
	"context"
	"errors"

	"github.com/edlanelima/system_sst/internal/domain/entity"
	"github.com/edlanelima/system_sst/internal/infraestructure/repository"
	"github.com/edlanelima/system_sst/internal/interfaces/dto"
)

type User struct {
	repo repository.User
}

func NewUser(repo repository.User) User {
	return User{repo: repo}
}

func (u *User) CreateUser(ctx context.Context, input dto.User) error {
	role, err := determineRole(input)
	if err != nil {
		return err
	}

	var user = entity.UserByRole{
		FullName: input.FullName,
		Email:    input.Email,
		Password: input.Password,
		Role:     role,
	}

	return u.repo.CreateUser(ctx, user)
}

func determineRole(input dto.User) (entity.Role, error) {
	switch {
	case input.IsSstRole:
		return entity.SST, nil
	case input.IsAdminRole:
		return entity.Admin, nil
	default:
		return "", errors.New("role could not be defined")
	}
}
