package main

import (
	"my_documents_south_backend/internal/core/repositories"
	"my_documents_south_backend/internal/core/services"
	"my_documents_south_backend/internal/repository/postgres"
	mdsHttp "my_documents_south_backend/internal/transport/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	db := postgres.New()
	defer db.Close()

	repository := repositories.NewRepositories(db)

	services := services.NewService(repository)

	httpHandler := mdsHttp.NewHttpHander(services)
	httpHandler.Route(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
