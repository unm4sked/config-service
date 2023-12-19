package configuration

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	CreateConfiguration()
	GetConfiguration(Id string)
	GetConfigurations()
	DeleteConfiguration(Id string)
	UpdateConfiguration(Id string)
}

type repository struct {
	Db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		Db: db,
	}
}

func (r *repository) CreateConfiguration() {
	// TODO
}

func (r *repository) GetConfiguration(Id string) {
	// TODO
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
