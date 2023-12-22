package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unm4sked/config-service/internal/api/response"
	"github.com/unm4sked/config-service/internal/configuration"
	"github.com/unm4sked/config-service/internal/entities"
)

func GetConfiguration(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		config, err := service.GetConfiguration(c.Params("id"))
		if err != nil {
			errMessages := []string{err.Error()}
			response := response.NewFailureResponseBody(errMessages)
			c.SendStatus(404)
			return c.JSON(response)
		}
		return c.JSON(response.NewSuccessResponseBody[entities.Configuration](config))
	}
}

func GetConfigurations(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		configs, err := service.GetConfigurations()
		if err != nil {
			errMessages := []string{err.Error()}
			response := response.NewFailureResponseBody(errMessages)
			c.SendStatus(400)
			return c.JSON(response)
		}
		return c.JSON(response.NewSuccessResponseBody[[]entities.Configuration](configs))
	}
}

func CreateConfiguration(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
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
		return c.JSON(response.NewSuccessResponseBody[entities.ConfigurationIdPayload](config))
	}
}

func DeleteConfiguration(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := service.DeleteConfiguration(id)
		if err != nil {
			errMessages := []string{err.Error()}
			response := response.NewFailureResponseBody(errMessages)
			c.SendStatus(422)
			return c.JSON(response)
		}
		return c.JSON(response.NewSuccessResponseBody[entities.ConfigurationIdPayload](entities.ConfigurationIdPayload{Id: id}))
	}
}

func UpdateConfiguration(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// id := c.Params("id")
		return c.JSON(fiber.Map{})
	}
}
