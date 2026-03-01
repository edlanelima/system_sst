package configs

import (
	"database/sql"
)

func NewMySQL() (*sql.DB, error) {

	dsn := "user:password@tcp(localhost:3306)/dbname?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
