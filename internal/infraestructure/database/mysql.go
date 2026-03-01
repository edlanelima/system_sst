package database

import (
	"context"
	"database/sql"
)

type IDatabase interface {
	SearchLevel(ctx context.Context, email, password string) (string, error)
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
