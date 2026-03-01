package service

import (
	"context"

	"github.com/edlanelima/system_sst/internal/domain/entity"
	"github.com/edlanelima/system_sst/internal/interfaces/dto"
)

type Login interface {
	LoginLevel(ctx context.Context, params dto.Login) (entity.Role, error)
}
