package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Test(c *fiber.Ctx) error {
	fmt.Println("Middleware")
	return c.Next()
}
