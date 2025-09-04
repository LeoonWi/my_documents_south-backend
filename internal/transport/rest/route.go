package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"my_documents_south_backend/internal/services"
	"my_documents_south_backend/internal/storage/postgres/repository"
	"time"
)

func Setup(db *sqlx.DB, app *fiber.App) {
	publicRouter := app.Group("")
	roleRoute(db, publicRouter)
	serviceRoute(db, publicRouter)
	tariffRoute(db, publicRouter)
}

func roleRoute(db *sqlx.DB, group fiber.Router) {
	repo := repository.NewRoleRepository(db)
	service := services.NewRoleService(repo, 10*time.Second)
	handler := NewRoleHandler(service)

	tag := group.Group("/roles")
	tag.Post("", handler.createRole)
	tag.Get("", handler.getRoles)
	tag.Get("/:id", handler.getRoleById)
	tag.Put("/:id", handler.updateRole)
	tag.Delete("/:id", handler.deleteRole)
}

func serviceRoute(db *sqlx.DB, group fiber.Router) {
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

func tariffRoute(db *sqlx.DB, group fiber.Router) {
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
