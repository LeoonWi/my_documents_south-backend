package middleware

import (
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTGenerate(userID int64, expiresIn time.Duration) (string, error) {
	// Создаем токен
	token := jwt.New(jwt.SigningMethodHS256)

	// Заполняем Payload
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(expiresIn).Unix()

	// Подписываем токен
	return token.SignedString([]byte("SECRET"))
}

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte("SECRET")},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
