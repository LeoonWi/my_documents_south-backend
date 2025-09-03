package http

import (
	"errors"
	"my_documents_south_backend/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (h *HttpHander) createRole(c *fiber.Ctx) error {
	role := &model.Role{}

	if err := c.BodyParser(role); err != nil {
		res := model.NewErrorResponse(errors.New(InvalidBody), c.Path()).Log()
		return c.Status(fiber.StatusUnprocessableEntity).JSON(res)
	}

	role, err := h.service.Role().CreateRole(c.Context(), role.Name)
	if err != nil {
		res := model.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusConflict).JSON(res)
	}

	return c.JSON(role)
}

func (h *HttpHander) getAllRole(c *fiber.Ctx) error {
	return c.JSON(h.service.Role().GetRoles(c.Context()))
}

func (h *HttpHander) getRoleById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		res := model.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	role, err := h.service.Role().GetRoleById(c.Context(), id)
	if err != nil {
		res := model.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusNotFound).JSON(res)
	}

	return c.JSON(role)
}

func (h *HttpHander) updateRole(c *fiber.Ctx) error {
	return nil
}

func (h *HttpHander) deleteRole(c *fiber.Ctx) error {
	return nil
}
