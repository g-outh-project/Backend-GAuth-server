package server

import (
	"log"

	"github.com/Backend-GAuth-server/server/middleware"
	v1 "github.com/Backend-GAuth-server/server/v1"
	auth "github.com/Backend-GAuth-server/server/v1/auth"
	"github.com/Backend-GAuth-server/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	app := fiber.New()
	file := utils.OpenLogger()

	defer file.Close()

	app.Use(limiter.New(utils.Limiter()))
	app.Use(logger.New(utils.Logger(file)))

	v1Router := app.Group("/api", middleware.JSONMiddleware)
	v1Router.Get("/life", v1.Life)

	authRouter := v1Router.Group("/auth")
	authRouter.Get("/login", auth.Login)
	authRouter.Get("/signup", auth.Signup)

	testRouter := v1Router.Group("/test", middleware.AuthMiddleware)
	testRouter.Get("/test", v1.Life)
	testRouter.Get("/shutdown", v1.Shutdown)
	log.Fatal(app.Listen(":8080"))
}
