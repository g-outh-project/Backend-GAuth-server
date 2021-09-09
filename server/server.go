package server

import (
	"log"
	"os"

	"github.com/Backend-GAuth-server/server/middleware"
	v1 "github.com/Backend-GAuth-server/server/v1"
	auth "github.com/Backend-GAuth-server/server/v1/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	app := fiber.New()
	file, err := os.OpenFile("./123.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app.Use(limiter.New())
	app.Use(logger.New(logger.Config{
		Format:     "${blue} [${time}] ${status} - ${method} ${path} ${latency}\n",
		TimeFormat: "15:04:03",
		TimeZone:   "Asia/Seoul",
	}))

	v1Router := app.Group("/api", middleware.JSONMiddleware)
	v1Router.Get("/life", v1.Life)

	authRouter := v1Router.Group("/auth")
	authRouter.Get("/login", auth.Login)
	authRouter.Get("/signup", auth.Signup)

	testRouter := v1Router.Group("/test", middleware.AuthMiddleware)
	testRouter.Get("/test", v1.Life)
	log.Fatal(app.Listen(":8080"))
}
