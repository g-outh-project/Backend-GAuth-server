package middleware

import (
	"fmt"

	"github.com/Backend-GAuth-server/method"
	"github.com/Backend-GAuth-server/utils"
	"github.com/gofiber/fiber/v2"
)

func JSONMiddleware(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	return c.Next()
}

func AuthMiddleware(c *fiber.Ctx) error {

	jwt, err := utils.GetTokenString(c)
	utils.HandleErr(err)

	_, user, err := utils.ValidateToken(string(jwt))

	fmt.Println(user.Id)
	data := method.SelectUserById(user.Id)

	fmt.Println(data.Name)
	if err != nil {
		return c.SendStatus(401)
	}
	return c.Next()
}
