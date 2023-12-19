package configuration

type Service interface {
	CreateConfiguration()
	GetConfigutation(Id string)
	GetConfigurations()
	UpdateConfiguration(Id string)
	DeleteConfiguration(Id string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateConfiguration() {
	s.repository.CreateConfiguration()
}
func (s *service) GetConfigutation(Id string) {
	s.repository.GetConfiguration(Id)
}
func (s *service) GetConfigurations() {
	s.repository.GetConfigurations()
}
func (s *service) UpdateConfiguration(Id string) {
	s.repository.UpdateConfiguration(Id)
}
func (s *service) DeleteConfiguration(Id string) {
	s.repository.DeleteConfiguration(Id)
}
