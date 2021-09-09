package v1

import "github.com/gofiber/fiber/v2"

type test struct {
	data string
}

func Life(c *fiber.Ctx) error {

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"posts":   "posts",
	})
}
