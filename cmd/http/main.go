package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/unm4sked/config-service/internal/api/routes"
	"github.com/unm4sked/config-service/internal/configuration"
	"github.com/unm4sked/config-service/internal/storage"
)

func main() {
	err := godotenv.Load("../../.env.dist")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgres, err := storage.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", postgres)

	configurationService := configuration.NewService(configuration.NewRepository(postgres.Db))

	app := fiber.New()
	api := app.Group("/api").Group("/v1")

	routes.ConfigurationRouter(api, configurationService)
	routes.SchemasRouter(api)

	log.Fatal(app.Listen(":3000"))
}
