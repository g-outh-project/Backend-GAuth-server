package auth

import (
	"fmt"

	"github.com/Backend-GAuth-server/dto"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var req dto.LoginReq
	c.Body()
	fmt.Println(req)
	return c.JSON(req)
}

func Signup(c *fiber.Ctx) error {
	var req dto.LoginReq
	fmt.Println(req)
	return c.JSON(req)
}
