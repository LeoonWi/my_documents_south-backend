package http

import (
	"github.com/gofiber/fiber/v2"
)

func (h HttpHander) signup_client(c *fiber.Ctx) error {
	return nil
}

func (h HttpHander) signup_employee(c *fiber.Ctx) error {
	return nil
}

func (h HttpHander) signin(c *fiber.Ctx) error {
	return nil
	//type User struct {
	//	userID   int64
	//	password string
	//}
	//
	//var user User
	//var err error
	//
	//if err = c.BodyParser(&user); err != nil {
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//		"message": err.Error(),
	//	})
	//}
	//
	//accessToken, err := middleware.JWTGenerate(user.userID, time.Minute*15)
	//refreshToken, err := middleware.JWTGenerate(user.userID, time.Hour*24)
	//
	//return c.Status(fiber.StatusOK).JSON(fiber.Map{"accessToken": accessToken, "refreshToken": refreshToken})
}

//
//func refreshfunc(c *fiber.Ctx) error {
//	// Получаем refresh-токен из тела запроса
//	var data struct {
//		RefreshToken string `json:"refresh_token"`
//	}
//	if err := c.BodyParser(&data); err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Невалидный запрос"})
//	}
//
//	// Проверяем refresh-токен
//	token, err := jwt.Parse(data.RefreshToken, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, jwt.ErrSignatureInvalid
//		}
//		return []byte("SECRET"), nil
//	})
//
//	if err != nil || !token.Valid {
//		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Невалидный refresh-токен"})
//	}
//
//	// Извлекаем данные из refresh-токена
//	claims := token.Claims.(jwt.MapClaims)
//	userID := claims["sub"].(string)
//	name := claims["name"].(string)
//
//	// Генерируем новый access-токен
//	newAccessToken, err := generateToken(userID, name, accessTokenSecret, time.Minute*15)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось сгенерировать новый access-токен"})
//	}
//
//	// (Опционально) Генерируем новый refresh-токен
//	newRefreshToken, err := generateToken(userID, name, refreshTokenSecret, time.Hour*24*7)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось сгенерировать новый refresh-токен"})
//	}
//
//	return c.JSON(fiber.Map{
//		"access_token":  newAccessToken,
//		"refresh_token": newRefreshToken,
//	})
//}
