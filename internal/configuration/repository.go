package configuration

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/unm4sked/config-service/internal/entities"
)

const ConfigurationTableName = "configurations"

type Repository interface {
	CreateConfiguration(id string, name string) error
	GetConfiguration(Id string)
	GetConfigurations()
	DeleteConfiguration(Id string)
	UpdateConfiguration(Id string)
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

func (r *repository) GetConfiguration(Id string) (entities.Configuration, error) {
	rows, err := r.db.Query(`SELECT * FROM configurations WHERE id=$1 LIMIT 1`, Id)
	defer rows.Close()
	if err != nil {
		log.Printf("An error occured while executing query: %v\n", err)
		return nil, err
	}

}

func (r *repository) GetConfigurations() {
	fmt.Println("GetConfigurations invoked")
	// TODO
}

func (r *repository) UpdateConfiguration(Id string) {
	// TODO
}

func (r *repository) DeleteConfiguration(Id string) {
	// TODO
}
