package utils

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Limiter() limiter.Config {
	app := limiter.Config{
		Max:        10,
		Expiration: 30 * time.Second,
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
		Format:     "${blue} [${time}] ${status} - ${method} ${path} ${latency}\n",
		TimeFormat: "15:04:03",
		TimeZone:   "Asia/Seoul",
	}

	return app
}

func FileLogger(file *os.File) logger.Config {
	app := logger.Config{
		Format:     "${blue} [${time}] ${status} - ${method} ${path} ${latency}\n",
		TimeFormat: "15:04:03",
		TimeZone:   "Asia/Seoul",
		Output:     file,
	}

	return app
}

func OpenLogger() *os.File {
	file, err := os.OpenFile("./"+time.Now().Format("20060102")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		HandlePanic(err)
	}

	return file
}
