package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Test(c *fiber.Ctx) error {
	fmt.Println("Middleware")
	c.Response().Header.Set("Content-Type", "application/json")
	return c.Next()
}
