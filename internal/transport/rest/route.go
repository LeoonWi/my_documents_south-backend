package rest

import (
	"my_documents_south_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Setup(db *sqlx.DB, app *fiber.App) {
	publicRouter := app.Group("/pub")

	protectedRouter := app.Group("/prot")
	protectedRouter.Use(middleware.Protected())

	ServiceRoute(db, protectedRouter)
	AuthRouter(
		publicRouter,
		protectedRouter,
		UserRoute(
			db,
			publicRouter,
			protectedRouter,
			TariffRoute(db, publicRouter, protectedRouter),
		),
		EmployeeRoute(db, publicRouter, protectedRouter, RoleRoute(db, publicRouter, protectedRouter)),
	)
}
