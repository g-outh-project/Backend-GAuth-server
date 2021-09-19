package server

import (
	"log"

	"github.com/Backend-GAuth-server/db"
	"github.com/Backend-GAuth-server/server/middleware"
	v1 "github.com/Backend-GAuth-server/server/v1"
	auth "github.com/Backend-GAuth-server/server/v1/auth"
	"github.com/Backend-GAuth-server/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() *fiber.App {
	// Basic Setting of server
	app := fiber.New()
	file := utils.OpenLogger()

	// Start and connect to DB
	db.Start()

	// Release resource
	defer db.CloseDB()
	defer file.Close()

	// Middleware setting
	app.Use(limiter.New(utils.Limiter()))
	app.Use(logger.New(utils.Logger(file)))

	// Routing
	v1Router := app.Group("/api", middleware.JSONMiddleware)
	v1Router.Get("/life", v1.Life)
	v1Router.Get("/refresh", auth.RefreshToken)

	authRouter := v1Router.Group("/auth")
	authRouter.Get("/login", auth.Login)
	authRouter.Get("/signup", auth.Signup)

	testRouter := v1Router.Group("/test", middleware.AuthMiddleware)
	testRouter.Get("/test", v1.Life)
	testRouter.Get("/shutdown", v1.Shutdown)
	log.Fatal(app.Listen(":8080"))

	return app
}
