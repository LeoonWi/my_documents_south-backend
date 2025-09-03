package docs

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func Init() fiber.Handler {
	return swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.yaml",
		Path:     "docs/swagger",
		Title:    "Swagger My Documents South API",
	})
}
