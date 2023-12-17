package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/unm4sked/config-service/internal/storage"
)

func main() {
	err := godotenv.Load("../../.env.dist")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	store, err := storage.NewConfigurationRepository()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", store)

	app := fiber.New()
	api := app.Group("/api").Group("/v1")

	api.Get("configurations", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	log.Fatal(app.Listen(":3000"))
}
