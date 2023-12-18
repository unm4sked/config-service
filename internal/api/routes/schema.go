package routes

import "github.com/gofiber/fiber/v2"

func SchemasRouter(app fiber.Router) {
	app.Get("/schemas", func(c *fiber.Ctx) error {
		return c.SendString("schemas")
	})
}
