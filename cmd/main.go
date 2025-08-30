package main

import (
	"github.com/jmoiron/sqlx"
	"log"
	"my_documents_south_backend/internal/core/repositories"
	"my_documents_south_backend/internal/core/services"
	mdsHttp "my_documents_south_backend/internal/transport/http"
	"my_documents_south_backend/pkg/storage/postgres"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	db := postgres.New()
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Panic(err)
		}
	}(db)

	repository := repositories.NewRepositories(db)

	services := services.NewService(repository)

	httpHandler := mdsHttp.NewHttpHander(services)
	httpHandler.Route(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
