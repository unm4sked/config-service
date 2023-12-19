package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresConnection struct {
	Db *sql.DB
}

func NewPostgresConnection() (*PostgresConnection, error) {
	connStr := "user=postgres dbname=app sslmode=disable password=password"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresConnection{
		Db: db,
	}, nil
}
