package utils

import (
	"crypto/sha512"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/savsgio/go-logger/v2"
)

// Handle critical error with throw panic
func HandlePanic(err error) {
	if err != nil {
		logger.Error(err)
		log.Panic(err) // exit
	}
}

// Handle error
func HandleErr(err error) {
	if err != nil {
		logger.Error(err)
	}
}

// Hash payload & return hash string
func Hash(payload interface{}) string {
	s := fmt.Sprintf("%v", payload)
	hash := sha512.Sum512([]byte(s))
	return fmt.Sprintf("%x", hash)
}

func GetSecretKey() string {
	// jwt Key
	var jwtConfig map[string]string
	jwtConfig, err := godotenv.Read()
	HandleErr(err)

	var jwtSignKey = jwtConfig["JWT_SECRET"]
	return jwtSignKey
}

func GetClientId(c *fiber.Ctx) string {
	clientId := c.Request().Header.Peek("clientId")
	cid := string(clientId[:])
	fmt.Println("--------clientId--------")
	fmt.Println(cid)
	return cid
}
