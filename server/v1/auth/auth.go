package auth

import (
	"fmt"

	"github.com/Backend-GAuth-server/dto"
	"github.com/Backend-GAuth-server/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var req dto.LoginReq
	utils.ByteToObj(c.Body(), &req)
	fmt.Println(req)
	utils.MarshalAndRes(200, req, c)
	return nil
}

func Signup(c *fiber.Ctx) error {
	var req dto.SignupReq
	utils.ByteToObj(c.Body(), &req)
	fmt.Println(req)
	c.SendStatus(201)
	return nil
}
