package repository

import (
	"context"

	"github.com/edlanelima/system_sst/internal/domain/entity"
	"github.com/edlanelima/system_sst/internal/infraestructure/database"
)

type Login struct {
	db database.IDatabase
}

func NewRepoLogin(db database.IDatabase) Login {
	return Login{db: db}
}

func (l *Login) SearchLevel(ctx context.Context, email, password string) (entity.Role, error) {
	roleStr, err := l.db.SearchLevel(ctx, email, password)
	if err != nil {
		return "", err
	}

	return entity.Role(roleStr), nil
}
