package database

import (
	"context"
	"errors"

	"github.com/edlanelima/system_sst/internal/domain/entity"
)

type DatabaseLocal struct{}

func NewMockDatabase() IDatabase {
	return &DatabaseLocal{}
}

func (d *DatabaseLocal) SearchLevel(
	ctx context.Context,
	email, password string,
) (string, error) {

	return mockBase(email, password)

}

func (d *DatabaseLocal) Insert(ctx context.Context, input entity.UserByRole) error {
	return nil
}

func mockBase(email, password string) (string, error) {
	if email == "admin@admin.com" && password == "admin" {
		return string(entity.Admin), nil
	}

	if email == "sst@sst.com" && password == "sst" {
		return string(entity.SST), nil
	}

	if email == "user@user.com" && password == "user" {
		return string(entity.User), nil
	}

	return "", errors.New("not found")
}
