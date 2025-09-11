package rest

import (
	"errors"
	"my_documents_south_backend/internal/middleware"
	"my_documents_south_backend/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	userService     models.UserService
	employeeService models.EmployeeService
}

func NewAuthHander(userService models.UserService, employeeService models.EmployeeService) *Auth {
	return &Auth{employeeService: employeeService, userService: userService}
}

func (h *Auth) loginUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		res := models.NewErrorResponse(errors.New("invalid body"), c.Path()).Log()
		return c.Status(fiber.StatusUnprocessableEntity).JSON(res)
	}

	accessToken, refreshToken, err := h.userService.Login(c.Context(), &user)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusConflict).JSON(res)
	}

	return c.JSON(fiber.Map{"access_token": accessToken, "refresh_token": refreshToken})
}

func (h *Auth) refreshToken(c *fiber.Ctx) error {
	var newAccessToken string
	var newRefreshToken string

	// Получаем refresh token из тела запроса или cookie
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
			"data":    nil,
		})
	}

	// Парсим и проверяем refresh token
	token, err := jwt.Parse(request.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte("my_documents_south_jwt_super_secret_key_for_security"), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid or expired refresh token",
			"data":    nil,
		})
	}

	userID, ok := c.Locals("userID").(int64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user ID in token",
			"data":    nil,
		})
	}

	roleID, ok := c.Locals("roleID").(int)
	if !ok {
		newAccessToken, err = middleware.JWTGenerate(userID, nil, time.Hour)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to generate access token",
				"data":    nil,
			})
		}

		newRefreshToken, err = middleware.JWTGenerate(userID, nil, 7*24*time.Hour)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to generate refresh token",
				"data":    nil,
			})
		}
	} else {
		newAccessToken, err = middleware.JWTGenerate(userID, &roleID, time.Hour)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to generate access token",
				"data":    nil,
			})
		}

		newRefreshToken, err = middleware.JWTGenerate(userID, &roleID, 7*24*time.Hour)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to generate refresh token",
				"data":    nil,
			})
		}
	}

	return c.JSON(fiber.Map{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}

func (h *Auth) loginEmployee(c *fiber.Ctx) error {
	var employee models.Employee

	if err := c.BodyParser(&employee); err != nil {
		res := models.NewErrorResponse(errors.New("invalid body"), c.Path()).Log()
		return c.Status(fiber.StatusUnprocessableEntity).JSON(res)
	}

	accessToken, refreshToken, err := h.employeeService.Login(c.Context(), &employee)
	if err != nil {
		res := models.NewErrorResponse(err, c.Path()).Log()
		return c.Status(fiber.StatusConflict).JSON(res)
	}

	return c.JSON(fiber.Map{"access_token": accessToken, "refresh_token": refreshToken})
}

func AuthRouter(
	public fiber.Router,
	protected fiber.Router,
	userService *models.UserService,
	employeeService *models.EmployeeService,
) {
	handler := NewAuthHander(*userService, *employeeService)

	public.Post("/users/signin", handler.loginUser)
	public.Post("/employee/signin", handler.loginEmployee)
	protected.Post("/auth/refresh", handler.refreshToken)
}
