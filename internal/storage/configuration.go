package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
	cfg "github.com/unm4sked/config-service/internal/configuration"
)

type ConfigurationRepository interface {
	CreateConfiguration(cfg.Configuration) error
	GetConfiguration(string) (cfg.Configuration, error)
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
