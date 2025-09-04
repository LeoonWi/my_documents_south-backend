package rest

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"my_documents_south_backend/internal/models"
)

type ServiceHandler struct {
	service models.ServiceService
}

func NewServiceHandler(service models.ServiceService) *ServiceHandler {
	return &ServiceHandler{service: service}
}

func (h *ServiceHandler) createService(c *fiber.Ctx) error {
	service := &models.Service{}

	if err := c.BodyParser(service); err != nil {
		res := models.NewErrorResponse(errors.New("Некорректное тело запроса"), c.Path()).Log()
		return c.Status(fiber.StatusUnprocessableEntity).JSON(res)
	}

	service, err := h.service.Create(c.Context(), service.Name)

	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusConflict).JSON(res)
	}

	return c.JSON(service)
}

func (h *ServiceHandler) getServices(c *fiber.Ctx) error {
	return c.JSON(h.service.Get(c.Context()))
}

func (h *ServiceHandler) getServiceById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	service, err := h.service.GetById(c.Context(), id)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusNotFound).JSON(res)
	}

	return c.JSON(service)
}

func (h *ServiceHandler) updateService(c *fiber.Ctx) error {
	// TODO update service handler
	return nil
}

func (h *ServiceHandler) deleteService(c *fiber.Ctx) error {
	// TODO delete service handler
	return nil
}
