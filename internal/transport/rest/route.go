package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Setup(db *sqlx.DB, app *fiber.App) {
	publicRouter := app.Group("")
	ClientRoute(db, publicRouter)
	EmployeeRoute(db, publicRouter)
	RoleRoute(db, publicRouter)
	ServiceRoute(db, publicRouter)
	TariffRoute(db, publicRouter)
}
