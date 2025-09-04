package rest

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"my_documents_south_backend/internal/models"
)

type RoleHandler struct {
	service models.RoleService
}

func NewRoleHandler(roleService models.RoleService) *RoleHandler {
	return &RoleHandler{service: roleService}
}

func (h *RoleHandler) createRole(c *fiber.Ctx) error {
	role := &models.Role{}

	if err := c.BodyParser(role); err != nil {
		res := models.NewErrorResponse(errors.New("Некорретное тело запроса"), c.Path()).Log()
		return c.Status(fiber.StatusUnprocessableEntity).JSON(res)
	}

	role, err := h.service.Create(c.Context(), role.Name)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusConflict).JSON(res)
	}

	return c.JSON(role)
}

func (h *RoleHandler) getRoles(c *fiber.Ctx) error {
	return c.JSON(h.service.Get(c.Context()))
}

func (h *RoleHandler) getRoleById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	role, err := h.service.GetById(c.Context(), id)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusNotFound).JSON(res)
	}

	return c.JSON(role)
}

func (h *RoleHandler) updateRole(c *fiber.Ctx) error {
	// TODO update role handler
	return nil
}

func (h *RoleHandler) deleteRole(c *fiber.Ctx) error {
	// TODO delete role handler
	return nil
}
