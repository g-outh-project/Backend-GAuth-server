package auth

import (
	"fmt"

	"github.com/Backend-GAuth-server/dto"
	"github.com/Backend-GAuth-server/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var req dto.LoginReq
	var res dto.LoginRes
	utils.ByteToObj(c.Body(), &req)
	res.AccessToken = utils.AccessToken(req.Id, req.Password)
	res.RefreshToken = utils.RefreshToken(req.Id, req.Password)
	utils.MarshalAndRes(200, res, c)
	return nil
}

func Signup(c *fiber.Ctx) error {
	var req dto.SignupReq
	utils.ByteToObj(c.Body(), &req)
	fmt.Println(req)
	c.SendStatus(201)
	return nil
}
