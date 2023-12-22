package entities

type CreateConfigurationPayload struct {
	Name string `json:"name"`
}

type ConfigurationIdPayload struct {
	Id string `json:"id"`
}

type Configuration struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
