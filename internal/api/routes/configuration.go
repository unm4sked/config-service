package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unm4sked/config-service/internal/configuration"
)

func ConfigurationRouter(app fiber.Router, repository configuration.Service) {
	app.Get("/configurations", func(c *fiber.Ctx) error {
		repository.GetConfigurations()
		return c.SendString("configurations")
	})
}
