package service

import (
	"context"

	"github.com/edlanelima/system_sst/internal/interfaces/dto"
)

type User interface {
	CreateUser(ctx context.Context, input dto.User) error
}
