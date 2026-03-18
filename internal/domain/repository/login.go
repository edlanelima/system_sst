package repository

import (
	"context"

	"github.com/edlanelima/system_sst/internal/domain/entity"
)

type Login interface {
	SearchLevel(ctx context.Context, email, password string) (entity.Role, error)
}
