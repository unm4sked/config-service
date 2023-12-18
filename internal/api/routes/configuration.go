package routes

import "github.com/gofiber/fiber/v2"

func ConfigurationRouter(app fiber.Router) {
	app.Get("/configurations", func(c *fiber.Ctx) error {
		return c.SendString("configurations")
	})
}
