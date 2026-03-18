package repository

import (
	"context"

	"github.com/edlanelima/system_sst/internal/domain/entity"
)

type User interface {
	CreateUser(ctx context.Context, input entity.UserByRole) error
}
