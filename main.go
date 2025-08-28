package main

import (
	"my_documents_south_backend/pkg/repository/postgres"
	"my_documents_south_backend/pkg/service"
	hHttp "my_documents_south_backend/pkg/transport/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	db := postgres.New()
	defer db.Close()

	repository := postgres.NewRepository(db)

	services := service.NewService(repository)

	http_handler := hHttp.NewHttpHander(services)
	http_handler.Route(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
