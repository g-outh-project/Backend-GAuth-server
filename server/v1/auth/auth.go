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

	res.AccessToken = utils.AccessToken(userData, c) // input token secret key in second argv
	userData.HashedAccessToken = utils.Hash(res.AccessToken)
	res.RefreshToken = utils.RefreshToken(userData, c) // input token secret key in second argv

	return c.Status(fiber.StatusOK).JSON(res)
}

func Signup(c *fiber.Ctx) error {
	var req dto.SignupReq

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	req.Password = utils.Hash(req.Password)
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
	_, user, err := utils.ValidateToken(req.RefreshToken, c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	jwt, err := utils.GetTokenString(c)
	utils.HandleErr(err)

	hashedJWT := utils.Hash(string(jwt))

	if hashedJWT != user.HashedAccessToken {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "invalid RefreshToken",
		})
	}

	jwtSource := dto.JWTSource{
		Id:                user.Id,
		Name:              user.Name,
		School:            user.School,
		Nickname:          user.Nickname,
		HashedAccessToken: "",
	}
	accessToken := utils.AccessToken(jwtSource, c)
	jwtSource.HashedAccessToken = utils.Hash(accessToken)
	refreshToken := utils.RefreshToken(jwtSource, c)

	return c.Status(fiber.StatusOK).JSON(dto.RefreshRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func GenKey(c *fiber.Ctx) error {
	jwt, err := utils.GetTokenString(c)
	utils.HandleErr(err)

	clientId := utils.Hash(string(jwt) + time.Now().UTC().String())
	jwtSecret := utils.Hash(string(jwt) + time.Now().String())

	err = method.InsertClient(clientId, jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to generate client Key",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"clientId":  clientId,
		"jwtSecret": jwtSecret,
	})
}
