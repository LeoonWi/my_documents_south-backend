package main

import (
	hHttp "my_documents_south_backend/pkg/handler/http"
	"my_documents_south_backend/pkg/repository/postgres"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	db := postgres.New()
	defer db.Close()

	hHttp.Route(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}

}
