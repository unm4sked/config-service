package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unm4sked/config-service/internal/api/handlers"
	"github.com/unm4sked/config-service/internal/configuration"
)

func ConfigurationRouter(app fiber.Router, service configuration.Service) {
	app.Get("/configurations/:id", handlers.GetConfiguration(service))
	app.Get("/configurations", handlers.GetConfigurations(service))
	app.Post("/configurations", handlers.CreateConfiguration(service))
	app.Delete("/configurations/:id", handlers.DeleteConfiguration(service))
	app.Put("/configurations/:id", handlers.UpdateConfiguration(service))
}
