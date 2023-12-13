package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type ConfigurationRepository interface {
	CreateConfiguration(int) error
	GetConfiguration(string) string
}

type PostgresConfigurationRepository struct {
	db *sql.DB
}

func NewConfigurationRepository() (*PostgresConfigurationRepository, error) {
	connStr := "user=postgres dbname=app sslmode=disable password=password"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresConfigurationRepository{
		db: db,
	}, nil
}
