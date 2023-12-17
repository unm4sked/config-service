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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
