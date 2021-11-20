package middleware

import (
	"github.com/Backend-GAuth-server/method"
	"github.com/Backend-GAuth-server/utils"
	"github.com/gofiber/fiber/v2"
)

func JSONMiddleware(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	c.Response().Header.Set("Access-Control-Allow-Origin", "*")
	return c.Next()
}

func AuthMiddleware(c *fiber.Ctx) error {

	jwt, err := utils.GetTokenString(c)
	utils.HandleErr(err)

	_, user, err := utils.ValidateToken(string(jwt))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "expired token",
		})
	}

	_, err = method.SelectUserById(user.Id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err != nil {
		return c.SendStatus(401)
	}
	return c.Next()
}
