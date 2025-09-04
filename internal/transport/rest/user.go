package rest

import (
	"my_documents_south_backend/internal/models"
	"my_documents_south_backend/internal/services"
	"my_documents_south_backend/internal/storage/postgres/repository"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type UserHandler struct {
	userService models.UserService
}

func NewUserHandler(userService models.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) createUser(c *fiber.Ctx) error {
	// TODO create user handler
	return nil
}

func (h *UserHandler) getUsers(c *fiber.Ctx) error {
	// TODO get users handler
	return nil
}

func (h *UserHandler) getUserById(c *fiber.Ctx) error {
	// TODO get user by id
	return nil
}

func (h *UserHandler) deleteUser(c *fiber.Ctx) error {
	// TODO delete user
	return nil
}

func ClientRoute(db *sqlx.DB, group fiber.Router) {
	repo := repository.NewUserRepository(db)
	service := services.NewUserService(repo, 10*time.Second)
	handler := NewUserHandler(service)

	tag := group.Group("/users")
	tag.Post("/signup", handler.createUser)
	tag.Get("", handler.getUsers)
	tag.Get("/:id", handler.getUserById)
	tag.Delete("/:id", handler.deleteUser)
}
