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

type RequestHandler struct {
	requestService models.RequestService
}

func NewRequestHandler(reqService models.RequestService) *RequestHandler {
	return &RequestHandler{reqService}
}

func (h *RequestHandler) createRequest(c *fiber.Ctx) error {
	var req models.Request

	if err := c.BodyParser(&req); err != nil {
		res := models.NewErrorResponse(errors.New("invalid body"), c.Path()).Log()
		return c.Status(fiber.StatusUnprocessableEntity).JSON(res)
	}

	err := h.requestService.Create(c.Context(), &req)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusConflict).JSON(res)
	}

	return c.Status(fiber.StatusCreated).JSON(req)
}

func (h *RequestHandler) getRequest(c *fiber.Ctx) error {
	return c.JSON(h.requestService.Get(c.Context()))
}

func (h *RequestHandler) getRequestById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	user, err := h.requestService.GetById(c.Context(), id)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusNotFound).JSON(res)
	}

	return c.JSON(user)
}

func (h *RequestHandler) getRequestsWithFilter(c *fiber.Ctx) error {
	filter := models.Request{}

	if ownerId := c.QueryInt("owner_id", 0); ownerId != 0 {
		filter.OwnerId = int64(ownerId)
	}
	if serviceId := c.QueryInt("service_id", 0); serviceId != 0 {
		filter.ServiceId = serviceId
	}
	if employeeId := c.QueryInt("employee_id", 0); employeeId != 0 {
		filter.EmployeeId = int64(employeeId)
	}

	if desiredAt := c.Query("desired_at"); desiredAt != "" {
		if t, err := time.Parse(time.RFC3339, desiredAt); err == nil {
			filter.DesiredAt = t
		}
	}

	requests, err := h.requestService.GetWithFilter(c.Context(), filter)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	return c.JSON(requests)
}

func (h *RequestHandler) deleteRequest(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		res := models.NewErrorResponse(errors.New("invalid id"), c.Path()).Log()
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	err = h.requestService.Delete(c.Context(), id)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusConflict).JSON(res)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func RequestRoute(db *sqlx.DB, group fiber.Router) {
	repo := repository.NewRequestRepository(db)
	service := services.NewRequestService(repo, 10*time.Second)
	handler := NewRequestHandler(service)

	tag := group.Group("/request")
	tag.Post("", handler.createRequest)
	tag.Get("", handler.getRequestsWithFilter)
	tag.Get("/:id", handler.getRequestById)
	tag.Delete("/:id", handler.deleteRequest)
}
