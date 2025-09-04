package rest

import (
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
	// TODO create employee handler
	return nil
}

func (h *EmployeeHandler) getEmployee(c *fiber.Ctx) error {
	// TODO get list employee handler
	return nil
}

func (h *EmployeeHandler) getEmployeeById(c *fiber.Ctx) error {
	// TODO get employee by id handler
	return nil
}

func (h *EmployeeHandler) deleteEmployee(c *fiber.Ctx) error {
	// TODO delete employee by id handler
	return nil
}

func EmployeeRoute(db *sqlx.DB, group fiber.Router) {
	repo := repository.NewEmployeeRepository(db)
	service := services.NewEmployeeService(repo, 10*time.Second)
	handler := NewEmployeeHandler(service)

	tag := group.Group("/employee")
	tag.Post("", handler.createEmployee)
	tag.Get("", handler.getEmployee)
	tag.Get("/:id", handler.getEmployeeById)
	tag.Delete("/:id", handler.deleteEmployee)
}
