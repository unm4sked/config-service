package entities

type CreateConfigurationPayload struct {
	Name string `json:"name"`
}

type ConfigurationCreatedPayload struct {
	Id string `json:"id"`
}

type Configuration struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
