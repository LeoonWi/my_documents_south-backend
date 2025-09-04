package rest

import (
	"errors"
	"my_documents_south_backend/internal/models"
	"my_documents_south_backend/internal/services"
	"my_documents_south_backend/internal/storage/postgres/repository"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type TariffHandler struct {
	service models.TariffService
}

func NewTariffHandler(tariffService models.TariffService) *TariffHandler {
	return &TariffHandler{tariffService}
}

func (h *TariffHandler) createTariff(c *fiber.Ctx) error {
	tariff := &models.Tariff{}
	if err := c.BodyParser(tariff); err != nil {
		res := models.NewErrorResponse(errors.New("Некорректное тело запроса"), c.Path()).Log()
		return c.Status(fiber.StatusUnprocessableEntity).JSON(res)
	}
	tariff, err := h.service.Create(c.Context(), tariff.Name)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusConflict).JSON(res)
	}
	return c.JSON(tariff)
}

func (h *TariffHandler) getTariffs(c *fiber.Ctx) error {
	return c.JSON(h.service.Get(c.Context()))
}

func (h *TariffHandler) getTariffById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	tariff, err := h.service.GetById(c.Context(), id)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusNotFound).JSON(res)
	}

	return c.JSON(tariff)
}

func (h *TariffHandler) updateTariff(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := c.BodyParser(&req); err != nil || req.Name == "" {
		res := models.NewErrorResponse(errors.New("invalid body"), c.Path()).Log()
		return c.Status(fiber.StatusUnprocessableEntity).JSON(res)
	}

	tariff, err := h.service.Update(c.Context(), id, req.Name)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusNotFound).JSON(res)
	}

	return c.JSON(tariff)
}

func (h *TariffHandler) deleteTariff(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		res := models.NewErrorResponse(errors.New("invalid id"), c.Path()).Log()
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	err = h.service.Delete(c.Context(), id)
	if err != nil {
		status := fiber.StatusInternalServerError
		if err.Error() == "tariff not found" {
			status = fiber.StatusNotFound
		}
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(status).JSON(res)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}

func TariffRoute(db *sqlx.DB, group fiber.Router) {
	repo := repository.NewTariffRepository(db)
	service := services.NewTariffService(repo, 10*time.Second)
	handler := NewTariffHandler(service)

	tag := group.Group("/tariffs")
	tag.Post("", handler.createTariff)
	tag.Get("", handler.getTariffs)
	tag.Get("/:id", handler.getTariffById)
	tag.Put("/:id", handler.updateTariff)
	tag.Delete("/:id", handler.deleteTariff)
}
