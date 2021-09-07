package server

import (
	"log"

	"github.com/Backend-GAuth-server/server/middleware"
	v1 "github.com/Backend-GAuth-server/server/v1"
	auth "github.com/Backend-GAuth-server/server/v1/auth"
	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()

	app.Group("/api")

	v1Router := app.Group("/api", middleware.Test)
	v1Router.Get("/life", v1.Life)

	authRouter := v1Router.Group("/auth", middleware.Test)
	authRouter.Get("/login", auth.Login)
	authRouter.Get("/signup", auth.Signup)

	log.Fatal(app.Listen(":8080"))
}
