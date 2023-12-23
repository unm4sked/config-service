package configuration

import (
	"database/sql"
	"log"

	"github.com/unm4sked/config-service/internal/entities"
)

const ConfigurationTableName = "configurations"

type Repository interface {
	CreateConfiguration(id string, name string) error
	GetConfiguration(id string) (entities.Configuration, error)
	GetConfigurations() ([]entities.Configuration, error)
	DeleteConfiguration(id string) error
	UpdateConfiguration(config entities.Configuration) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateConfiguration(id string, name string) error {
	_, err := r.db.Exec(`INSERT INTO configurations ("id", "name") VALUES ($1,$2)`, id, name)
	if err != nil {
		log.Printf("An error occured while executing query: %v\n", err)
		return err
	}

	return nil
}

func (r *repository) GetConfiguration(id string) (entities.Configuration, error) {
	row := r.db.QueryRow(`SELECT * FROM configurations WHERE id=$1 LIMIT 1`, id)
	var configuration = entities.Configuration{}
	err := row.Scan(&configuration.Id, &configuration.Name)
	if err != nil {
		log.Printf("An error occured while executing query: %v\n", err)
		return entities.Configuration{}, err
	}

	return configuration, nil
}

func (r *repository) GetConfigurations() ([]entities.Configuration, error) {
	rows, err := r.db.Query(`SELECT * FROM configurations`)
	configs := []entities.Configuration{}
	if err != nil {
		return configs, err
	}
	defer rows.Close()

	for rows.Next() {
		config := entities.Configuration{}
		rows.Scan(&config.Id, &config.Name)
		configs = append(configs, config)
	}

	return configs, nil
}

func (r *repository) UpdateConfiguration(config entities.Configuration) error {

	_, err := r.db.Exec(`UPDATE configurations SET name=$1 WHERE id=$2`, config.Name, config.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteConfiguration(Id string) error {
	_, err := r.db.Exec("DELETE FROM configurations WHERE id=$1", Id)
	if err != nil {
		return err
	}
	return nil
}
