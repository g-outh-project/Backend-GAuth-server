package middleware

import (
	"fmt"

	"github.com/Backend-GAuth-server/utils"
	"github.com/gofiber/fiber/v2"
)

func JSONMiddleware(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	return c.Next()
}

func AuthMiddleware(c *fiber.Ctx) error {
	jwt, _ := utils.GetTokenString(c)
	fmt.Println(jwt)
	return c.Next()
}
