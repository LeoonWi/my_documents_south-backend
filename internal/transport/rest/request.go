package rest

import (
	"errors"
	"my_documents_south_backend/internal/models"
	"my_documents_south_backend/internal/services"
	"my_documents_south_backend/internal/storage/postgres/repository"
	"strconv"
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

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *RequestHandler) getRequestsWithFilter(c *fiber.Ctx) error {
	filter := models.Request{}

	if ownerIdStr := c.Query("owner_id"); ownerIdStr != "" {
		if ownerId, err := strconv.ParseInt(ownerIdStr, 10, 64); err == nil {
			filter.OwnerId = ownerId
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid owner_id"})
		}
	}
	if serviceIdStr := c.Query("service_id"); serviceIdStr != "" {
		if serviceId, err := strconv.ParseInt(serviceIdStr, 10, 32); err == nil {
			filter.ServiceId = int(serviceId)
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid service_id"})
		}
	}

	if statusStr := c.Query("status"); statusStr != "" {
		if status, err := strconv.ParseInt(statusStr, 10, 16); err == nil {
			filter.Status = int16(status)
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid status"})
		}
	}

	if employeeIdStr := c.Query("employee_id"); employeeIdStr != "" {
		if employeeId, err := strconv.ParseInt(employeeIdStr, 10, 64); err == nil {
			filter.EmployeeId = employeeId
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid employee_id"})
		}
	}

	if desiredAt := c.Query("desired_at"); desiredAt != "" {
		if t, err := time.Parse(time.RFC3339, desiredAt); err == nil {
			filter.DesiredAt = t
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid desired_at"})
		}
	}

	requests, err := h.requestService.GetWithFilter(c.Context(), filter)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	return c.Status(fiber.StatusOK).JSON(requests)
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

func RequestRoute(db *sqlx.DB, protected fiber.Router, user *models.UserService, employee *models.EmployeeService) {
	repo := repository.NewRequestRepository(db)
	service := services.NewRequestService(repo, 10*time.Second)

	handler := NewRequestHandler(service)

	//tag := group.Group("/request")
	protected.Post("/request", handler.createRequest)
	protected.Get("/request", handler.getRequestsWithFilter)
	protected.Get("/request/:id", handler.getRequestById)
	//protected.Patch("/request/:id/:employee_id", handler.deleteRequest)
	//protected.Patch("/request/:id/:status", handler.deleteRequest)
	protected.Delete("/request/:id", handler.deleteRequest)
}
