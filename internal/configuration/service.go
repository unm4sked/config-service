package configuration

import (
	"github.com/google/uuid"
	"github.com/unm4sked/config-service/internal/entities"
)

type Service interface {
	CreateConfiguration(name string) (entities.ConfigurationIdPayload, error)
	GetConfiguration(id string) (entities.Configuration, error)
	GetConfigurations() ([]entities.Configuration, error)
	UpdateConfiguration(config entities.Configuration) error
	DeleteConfiguration(id string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateConfiguration(name string) (entities.ConfigurationIdPayload, error) {
	id := uuid.NewString()
	err := s.repository.CreateConfiguration(id, name)
	configurationCreated := entities.ConfigurationIdPayload{}
	if err != nil {
		return configurationCreated, err
	}
	return entities.ConfigurationIdPayload{Id: id}, nil
}
func (s *service) GetConfiguration(id string) (entities.Configuration, error) {
	return s.repository.GetConfiguration(id)
}
func (s *service) GetConfigurations() ([]entities.Configuration, error) {
	return s.repository.GetConfigurations()
}
func (s *service) UpdateConfiguration(config entities.Configuration) error {
	return s.repository.UpdateConfiguration(config)
}
func (s *service) DeleteConfiguration(Id string) error {
	return s.repository.DeleteConfiguration(Id)
}
