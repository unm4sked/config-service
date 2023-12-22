package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unm4sked/config-service/internal/configuration"
	"github.com/unm4sked/config-service/internal/entities"
)

func ConfigurationRouter(app fiber.Router, service configuration.Service) {
	app.Get("/configurations/:id", func(c *fiber.Ctx) error {
		config, err := service.GetConfiguration(c.Params("id"))
		if err != nil {
			return fiber.ErrNotFound
		}
		return c.JSON(config)
	})
	app.Get("/configurations", func(c *fiber.Ctx) error {
		configs, err := service.GetConfigurations()
		if err != nil {
			return fiber.ErrInternalServerError
		}
		return c.JSON(configs)
	})
	app.Post(("/configurations"), func(c *fiber.Ctx) error {
		payload := new(entities.CreateConfigurationPayload)
		if err := c.BodyParser(payload); err != nil {
			return err
		}
		id, err := service.CreateConfiguration(payload.Name)
		if err != nil {
			return err
		}

		return c.JSON(&entities.ConfigurationCreatedPayload{Id: id})
	})
}
