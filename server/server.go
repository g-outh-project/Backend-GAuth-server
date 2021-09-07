package server

import (
	"log"
	"os"

	"github.com/Backend-GAuth-server/server/middleware"
	v1 "github.com/Backend-GAuth-server/server/v1"
	auth "github.com/Backend-GAuth-server/server/v1/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	app := fiber.New()
	file, err := os.OpenFile("./123.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app.Use(logger.New(logger.Config{
		Output: file,
	}))
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		// Format:     "${pid} ${status} - ${method} ${path}\n",
		// TimeFormat: "02-Jan-2006",
		// TimeZone:   "America/New_York",
	}))

	app.Group("/api")

	v1Router := app.Group("/api", middleware.Test)
	v1Router.Get("/life", v1.Life)

	authRouter := v1Router.Group("/auth", middleware.Test)
	authRouter.Get("/login", auth.Login)
	authRouter.Get("/signup", auth.Signup)

	log.Fatal(app.Listen(":8080"))
}
