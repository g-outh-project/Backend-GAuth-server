package auth

import "github.com/gofiber/fiber/v2"

type test struct {
	Data string `json:"data"`
}

func Login(c *fiber.Ctx) error {
	res := test{Data: "auth"}
	return c.JSON(res)
}

func Signup(c *fiber.Ctx) error {
	res := test{Data: "signup"}
	return c.JSON(res)
}
