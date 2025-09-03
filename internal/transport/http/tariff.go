package http

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"my_documents_south_backend/internal/model"
)

func (h *HttpHander) createTariff(c *fiber.Ctx) error {
	tariff := &model.Tariff{}
	if err := c.BodyParser(tariff); err != nil {
		res := model.NewErrorResponse(errors.New(InvalidBody), c.Path()).Log()
		return c.Status(fiber.StatusUnprocessableEntity).JSON(res)
	}
	tariff, err := h.service.Tariff().CreateTariff(c.Context(), tariff.Name)
	if err != nil {
		res := model.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusConflict).JSON(res)
	}
	return c.JSON(tariff)
}

func (h *HttpHander) getAllTariff(c *fiber.Ctx) error {
	return c.JSON(h.service.Tariff().GetTariffs(c.Context()))
}

func (h *HttpHander) getTariffById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		res := model.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	tariff, err := h.service.Tariff().GetTariffByID(c.Context(), id)
	if err != nil {
		res := model.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusNotFound).JSON(res)
	}

	return c.JSON(tariff)
}

func (h *HttpHander) getTariffByName(c *fiber.Ctx) error {
	name := c.Params("name")
	if name == "" {
		res := model.NewErrorResponse(errors.New("name is required"), c.Path()).Log()
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	tariff, err := h.service.Tariff().GetTariffByName(c.Context(), name)
	if err != nil {
		res := model.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusNotFound).JSON(res)
	}

	return c.JSON(tariff)
}

func (h *HttpHander) updateTariff(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		res := model.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := c.BodyParser(&req); err != nil || req.Name == "" {
		res := model.NewErrorResponse(errors.New("invalid body"), c.Path()).Log()
		return c.Status(fiber.StatusUnprocessableEntity).JSON(res)
	}

	tariff, err := h.service.Tariff().UpdateTariff(c.Context(), id, req.Name)
	if err != nil {
		res := model.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusNotFound).JSON(res)
	}

	return c.JSON(tariff)
}

func (h *HttpHander) deleteTariff(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		res := model.NewErrorResponse(errors.New("invalid id"), c.Path()).Log()
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	err = h.service.Tariff().DeleteTariff(c.Context(), id)
	if err != nil {
		status := fiber.StatusInternalServerError
		if err.Error() == "tariff not found" {
			status = fiber.StatusNotFound
		}
		res := model.NewErrorResponse(err, c.Path()).Log()
		return c.Status(status).JSON(res)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "tariff deleted successfully",
		"id":      id,
	})
}
