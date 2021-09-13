package auth

import (
	"time"

	"github.com/Backend-GAuth-server/dto"
	"github.com/Backend-GAuth-server/method"
	"github.com/Backend-GAuth-server/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var req dto.LoginReq
	var res dto.LoginRes

	userData := dto.JWTSource{
		Uid:               "123",
		Id:                "123",
		Email:             "123",
		Name:              "123",
		School:            "123",
		Birth:             "123",
		Nickname:          "123",
		HashedAccessToken: "",
		CreatedAt:         time.Now().UTC(),
	}

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	res.AccessToken = utils.AccessToken(userData)
	userData.HashedAccessToken = utils.Hash(res.AccessToken)
	res.RefreshToken = utils.RefreshToken(userData)
	utils.MarshalAndRes(200, res, c)
	return nil
}

func Signup(c *fiber.Ctx) error {
	var req dto.SignupReq

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	method.InsertUser(req)

	c.SendStatus(201)
	return nil
}
