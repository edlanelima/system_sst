package repository

import (
	"context"

	"github.com/edlanelima/system_sst/internal/domain/entity"
	"github.com/edlanelima/system_sst/internal/infraestructure/database"
)

type User struct {
	db database.IDatabase
}

func NewRepoUser(db database.IDatabase) User {
	return User{db: db}
}

func (u *User) CreateUser(ctx context.Context, input entity.UserByRole) error {
	return u.db.Insert(ctx, input)
}
