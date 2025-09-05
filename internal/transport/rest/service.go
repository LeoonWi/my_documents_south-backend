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

	service, err := h.service.Update(c.Context(), id, req.Name)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusNotFound).JSON(res)
	}

	return c.JSON(service)
}

func (h *ServiceHandler) deleteService(c *fiber.Ctx) error {
	// TODO delete service handler
	return nil
}

func ServiceRoute(db *sqlx.DB, group fiber.Router) {
	repo := repository.NewServiceRepository(db)
	service := services.NewServiceService(repo, 10*time.Second)
	handler := NewServiceHandler(service)

	tag := group.Group("/services")
	tag.Post("", handler.createService)
	tag.Get("", handler.getServices)
	tag.Get("/:id", handler.getServiceById)
	tag.Put("/:id", handler.updateService)
	tag.Delete("/:id", handler.deleteService)
}
