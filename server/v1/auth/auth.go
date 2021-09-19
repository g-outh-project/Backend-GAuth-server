package auth

import (
	"fmt"

	"github.com/Backend-GAuth-server/dto"
	"github.com/Backend-GAuth-server/method"
	"github.com/Backend-GAuth-server/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var req dto.LoginReq
	var res dto.LoginRes

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := method.SelectUserById(req.Id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if user.Password != utils.Hash(req.Password) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "invalid password",
		})
	}

	userData := dto.JWTSource{
		Id:                user.Id,
		Name:              user.Name,
		Nickname:          user.Nickname,
		School:            user.School,
		HashedAccessToken: "",
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

	err = method.InsertUser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "same id is exist",
		})
	}

	c.SendStatus(201)
	return nil
}

func RefreshToken(c *fiber.Ctx) error {
	var req dto.RefreshReq

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	_, user, err := utils.ValidateToken(req.RefreshToken)

	jwt, err := utils.GetTokenString(c)
	utils.HandleErr(err)

	hashedJWT := utils.Hash(string(jwt))

	if hashedJWT != user.HashedAccessToken {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "invalid RefreshToken",
		})
	}

	fmt.Println(hashedJWT)
	fmt.Println(user.HashedAccessToken)
	return nil
}
