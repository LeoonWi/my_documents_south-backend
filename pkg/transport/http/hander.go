package http

import (
	"my_documents_south_backend/pkg/service"

	"github.com/gofiber/fiber/v2"
)

type HttpHander struct {
	service *service.Service
}

func NewHttpHander(service *service.Service) *HttpHander {
	return &HttpHander{service: service}
}

func (h HttpHander) Route(app *fiber.App) {
	// auth
	auth := app.Group("/auth")
	auth.Post("/signin", h.signin)
	//auth.POST("/refresh", refresh)

	client := app.Group("/client")
	client.Post("/signup", h.signup_client)

	employee := app.Group("/employee")
	employee.Post("/signup", h.signup_employee)

	// tariff
	tariff := app.Group("/tariff")
	tariff.Post("", h.createTariff)
	tariff.Get("", h.getAllTariff)
	tariff.Get("/:id", h.getTariffById)
	tariff.Put("/:id", h.updateTariff)
	tariff.Delete("/:id", h.deleteTariff)

	// role
	role := app.Group("/role")
	role.Post("", h.createRole)
	role.Get("", h.getAllRole)
	role.Get("/:id", h.getRoleById)
	role.Put("/:id", h.updateRole)
	role.Delete("/:id", h.deleteRole)

	// service
	svc := app.Group("/service")
	svc.Post("", h.createService)
	svc.Get("", h.getAllService)
	svc.Get("/:id", h.getServiceById)
	svc.Put("/:id", h.updateService)
	svc.Delete("/:id", h.deleteService)
}
