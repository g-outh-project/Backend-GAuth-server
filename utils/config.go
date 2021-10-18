package utils

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/utils"
)

func Limiter() limiter.Config {
	app := limiter.Config{
		Max:        2,
		Expiration: 1 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(&fiber.Map{
				"statusCode":   429,
				"errorMessage": "Too many Request",
			})
		},
	}

	return app
}

func ConsoleLogger() logger.Config {
	app := logger.Config{
		Format:     "${blue} [${time}] [${ip}] ${status} - ${method} ${path} ${latency}\n",
		TimeFormat: "15:04:03",
		TimeZone:   "Asia/Seoul",
	}

	return app
}

func FileLogger(file *os.File) logger.Config {
	app := logger.Config{
		Format:     "${blue} [${time}] [${ip}] ${status} - ${method} ${path} ${latency}\n",
		TimeFormat: "15:04:03",
		TimeZone:   "Asia/Seoul",
		Output:     file,
	}

	return app
}

func OpenLogger() *os.File {
	err := os.Mkdir("logs", 0777)
	HandleErr(err)

	file, err := os.OpenFile("./logs/"+time.Now().Format("20060102")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		HandlePanic(err)
	}

	return file
}

func Csrf() csrf.Config {
	app := csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
	}

	return app
}
