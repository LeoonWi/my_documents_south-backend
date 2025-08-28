package http

import "github.com/gofiber/fiber/v2"

func Route(app *fiber.App) {
	// auth
	auth := app.Group("/auth")
	auth.Post("/signin", signin)
	//auth.POST("/refresh", refresh)

	client := app.Group("/client")
	client.Post("/signup", signup_client)

	employee := app.Group("/employee")
	employee.Post("/signup", signup_employee)

	// tariff
	tariff := app.Group("/tariff")
	tariff.Post("", createTariff)
	tariff.Get("", getAllTariff)
	tariff.Get("/:name", getTariffByName)
	tariff.Put("/:name", updateTariff)
	tariff.Delete("/:name", deleteTariff)

	// role
	role := app.Group("/role")
	role.Post("", createRole)
	role.Get("", getAllRole)
	role.Get("/:name", getRoleByName)
	role.Put("/:name", updateRole)
	role.Delete("/:name", deleteRole)

	// service
	service := app.Group("/service")
	service.Post("", createService)
	service.Get("", getAllService)
	service.Get("/:id", getServiceById)
	service.Put("/:id", updateService)
	service.Delete("/:id", deleteService)
}
