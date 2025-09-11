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

type EmployeeHandler struct {
	employeeService models.EmployeeService
}

func NewEmployeeHandler(employeeService models.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{employeeService: employeeService}
}

func (h *EmployeeHandler) createEmployee(c *fiber.Ctx) error {
	var employee models.Employee

	if err := c.BodyParser(&employee); err != nil {
		res := models.NewErrorResponse(errors.New("invalid body"), c.Path()).Log()
		return c.Status(fiber.StatusUnprocessableEntity).JSON(res)
	}

	err := h.employeeService.Create(c.Context(), &employee)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusConflict).JSON(res)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *EmployeeHandler) getEmployee(c *fiber.Ctx) error {
	return c.JSON(h.employeeService.Get(c.Context()))
}

func (h *EmployeeHandler) getEmployeeById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	employee, err := h.employeeService.GetById(c.Context(), id)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusNotFound).JSON(res)
	}

	return c.JSON(employee)
}

func (h *EmployeeHandler) deleteEmployee(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		res := models.NewErrorResponse(errors.New("invalid id"), c.Path()).Log()
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	err = h.employeeService.Delete(c.Context(), id)
	if err != nil {
		status := fiber.StatusInternalServerError
		if err.Error() == "user not found" {
			status = fiber.StatusNotFound
		}
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(status).JSON(res)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}

func EmployeeRoute(db *sqlx.DB, public fiber.Router, protected fiber.Router, roleRepo *models.RoleRepository) *models.EmployeeService {
	repo := repository.NewEmployeeRepository(db)

	service := services.NewEmployeeService(repo, *roleRepo, 10*time.Second)
	handler := NewEmployeeHandler(service)

	public.Post("/employee/signup", handler.createEmployee)
	protected.Get("/employee", handler.getEmployee)
	protected.Get("/employee/:id", handler.getEmployeeById)
	protected.Delete("/employee/:id", handler.deleteEmployee)

	return &service
}
