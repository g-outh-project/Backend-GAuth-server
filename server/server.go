package server

import (
	"fmt"
	"log"

	"github.com/Backend-GAuth-server/db"
	"github.com/Backend-GAuth-server/server/middleware"
	v1 "github.com/Backend-GAuth-server/server/v1"
	auth "github.com/Backend-GAuth-server/server/v1/auth"
	"github.com/Backend-GAuth-server/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Start() {
	// Basic Setting of server
	app := fiber.New(fiber.Config{
		Prefork: false,
	})
	file := utils.OpenLogger()

	// Start and connect to DB
	db.Start()

	// Middleware setting
	app.Use(cors.New())
	app.Use(pprof.New())
	app.Use(recover.New())

	//app.Use(csrf.New(utils.Csrf()))
	app.Use(limiter.New(utils.Limiter()))
	app.Use(logger.New(utils.ConsoleLogger()))
	app.Use(logger.New(utils.FileLogger(file)))

	api := app.Group("/api")
	api.Get("/dashboard", monitor.New())

	// Routing
	v1Router := api.Group("/v1")
	v1Router.Get("/life", v1.Life)
	v1Router.Put("/refresh", auth.RefreshToken)

	authRouter := v1Router.Group("/auth")
	authRouter.Post("/login", auth.Login)
	authRouter.Post("/signup", auth.Signup)

	testRouter := v1Router.Group("/test", middleware.AuthMiddleware)
	testRouter.Get("/test", v1.Life)
	testRouter.Get("/shutdown", v1.Shutdown)

	defer file.Close()
	defer db.CloseDB()
	log.Fatal(app.Listen(":" + fmt.Sprint(8080)))
}
