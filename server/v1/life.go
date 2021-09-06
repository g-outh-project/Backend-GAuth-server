package v1

import "github.com/gofiber/fiber/v2"

type test struct {
	data string
}

func Life(c *fiber.Ctx) error {
	res := test{data: "hello"}
	return c.JSON(res)
}
