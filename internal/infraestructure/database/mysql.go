package database

import (
	"context"
	"database/sql"

	"github.com/edlanelima/system_sst/internal/domain/entity"
)

type IDatabase interface {
	SearchLevel(ctx context.Context, email, password string) (string, error)
	Insert(ctx context.Context, input entity.UserByRole) error
}

type Database struct {
	mysql *sql.DB
}

func NewDbMySQL(db *sql.DB) IDatabase {
	return &Database{mysql: db}
}

func (d *Database) SearchLevel(
	ctx context.Context,
	email, password string,
) (string, error) {
	query := `
        SELECT role
        FROM users
        WHERE email = ?
        AND password = ?
    `

	var role string

	row := d.mysql.QueryRowContext(ctx, query, email, password)

	err := row.Scan(&role)
	if err != nil {
		return "", err
	}

	return role, nil
}

func (d *Database) Insert(ctx context.Context, input entity.UserByRole) error {
	query := `
		INSERT INTO users (fullname, email, password, level)
		VALUES (?, ?, ?, ?)
	`

	_, err := d.mysql.ExecContext(ctx, query, input.FullName, input.Email, input.Password, input.Role)
	if err != nil {
		return err
	}

	return nil
}
