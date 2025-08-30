package main

import (
	"my_documents_south_backend/internal/core/facade"
	postgres2 "my_documents_south_backend/internal/storage/postgres"
	mdsHttp "my_documents_south_backend/internal/transport/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	db := postgres2.New()
	defer db.Close()

	repository := facade.NewRepository(db)

	services := facade.NewService(repository)

	http_handler := mdsHttp.NewHttpHander(services)
	http_handler.Route(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
