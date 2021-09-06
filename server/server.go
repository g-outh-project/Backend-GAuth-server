package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		msg := fmt.Sprintln("Hello")
		return c.SendString(msg)
	})

	log.Fatal(app.Listen(":8080"))
}
