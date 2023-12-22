package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unm4sked/config-service/internal/api/response"
	"github.com/unm4sked/config-service/internal/configuration"
	"github.com/unm4sked/config-service/internal/entities"
)

func ConfigurationRouter(app fiber.Router, service configuration.Service) {
	app.Get("/configurations/:id", func(c *fiber.Ctx) error {
		config, err := service.GetConfiguration(c.Params("id"))
		if err != nil {
			errMessages := []string{err.Error()}
			response := response.NewFailureResponseBody(errMessages)
			c.SendStatus(404)
			return c.JSON(response)
		}

		return c.JSON(response.NewSuccessResponseBody[entities.Configuration](config))
	})

	app.Get("/configurations", func(c *fiber.Ctx) error {
		configs, err := service.GetConfigurations()
		if err != nil {
			errMessages := []string{err.Error()}
			response := response.NewFailureResponseBody(errMessages)
			c.SendStatus(400)
			return c.JSON(response)
		}
		return c.JSON(response.NewSuccessResponseBody[[]entities.Configuration](configs))
	})

	app.Post(("/configurations"), func(c *fiber.Ctx) error {
		payload := new(entities.CreateConfigurationPayload)
		if err := c.BodyParser(payload); err != nil {
			errMessages := []string{err.Error()}
			response := response.NewFailureResponseBody(errMessages)
			c.SendStatus(400)
			return c.JSON(response)
		}
		config, err := service.CreateConfiguration(payload.Name)
		if err != nil {
			errMessages := []string{err.Error()}
			response := response.NewFailureResponseBody(errMessages)
			c.SendStatus(422)
			return c.JSON(response)
		}
		return c.JSON(response.NewSuccessResponseBody[entities.ConfigurationCreatedPayload](config))
	})
}
