package http

import (
	"errors"
	"my_documents_south_backend/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (h *HttpHander) createService(c *fiber.Ctx) error {
	service := &model.Service{}

	if err := c.BodyParser(service); err != nil {
		res := model.NewErrorResponse(errors.New(InvalidBody), c.Path()).Log()
		return c.Status(fiber.StatusUnprocessableEntity).JSON(res)
	}

	service, err := h.service.Service().CreateService(c.Context(), service.Name)

	if err != nil {
		res := model.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusConflict).JSON(res)
	}

	return c.JSON(service)
}

func (h *HttpHander) getAllService(c *fiber.Ctx) error {
	return c.JSON(h.service.Service().GetService(c.Context()))
}

func (h *HttpHander) getServiceById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		res := model.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	service, err := h.service.Service().GetServiceById(c.Context(), id)
	if err != nil {
		res := model.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusNotFound).JSON(res)
	}

	return c.JSON(service)
}

func (h *HttpHander) updateService(c *fiber.Ctx) error {
	return nil
}

func (h *HttpHander) deleteService(c *fiber.Ctx) error {
	return nil
}
