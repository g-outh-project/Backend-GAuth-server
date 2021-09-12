package v1

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

type test struct {
	data string
}

func Life(c *fiber.Ctx) error {

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"posts":   "posts",
	})
}

func Shutdown(c *fiber.Ctx) error {
	log.Panic()
	return errors.New("hello")
}
