package http

import "github.com/gofiber/fiber/v2"

func Route(app *fiber.App) {
	// auth
	auth := app.Group("/auth")
	auth.Post("/signin", signin)

	client := app.Group("/client")
	client.Post("/signup", signup_client)

	employee := app.Group("/employee")
	employee.Post("/signup", signup_employee)

	// tariff
	tariff := app.Group("/tariff")
	tariff.Post("", nil)
	tariff.Get("", nil)
	tariff.Get("/{name}", nil)
	tariff.Put("/{name}", nil)
	tariff.Delete("/{name}", nil)

	// role
	role := app.Group("/role")
	role.Post("", nil)
	role.Get("", nil)
	role.Get("/{name}", nil)
	role.Put("/{name}", nil)
	role.Delete("/{name}", nil)

	// service
	service := app.Group("/service")
	service.Post("", nil)
	service.Get("", nil)
	service.Get("/{id}", nil)
	service.Put("/{id}", nil)
	service.Delete("/{id}", nil)
}
