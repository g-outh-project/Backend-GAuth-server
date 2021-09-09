package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// test jwtSignKey
var testSignKey = []byte("TestForFasthttpWithJWT")

// Credential type
type userCredential struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GetToken function
func GetTokenString(c *fiber.Ctx) ([]byte, error) {
	// Get token from request token
	jwt := c.Request().Header.Peek("Authorization")

	// Token length validation
	if len(jwt) == 0 {
		c.SendStatus(401)
		return nil, errors.New("Token cannot found")
	}

	// Return token with type []byte
	return jwt, nil
}

// Generate accessToken
func AccessToken(id string, password string) string {
	// Generate Token object
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		Id:       id,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(), // 10 Mins
		},
	})

	// jwt Key
	jwtSignKey := []byte(GetSecretKey())

	// Sign token
	access, err := accessToken.SignedString(jwtSignKey)
	HandleErr(err)

	return access
}

// Generate refreshToken
func RefreshToken(id string, password string) string {
	// Generate Token object
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		Id:       id,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(336 * time.Hour).Unix(), // 14Days
		},
	})

	// jwt Key
	jwtSignKey := []byte(GetSecretKey())

	// Sign token
	refresh, err := refreshToken.SignedString(jwtSignKey)
	HandleErr(err)

	return refresh
}

// Validate token
func ValidateToken(requestToken string) (*jwt.Token, *userCredential, error) {
	// Generate Credential object
	user := &userCredential{}

	// jwt Key
	jwtSignKey := []byte(GetSecretKey())

	// Parse token and validate
	token, err := jwt.ParseWithClaims(requestToken, user, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})
	HandleErr(err)

	return token, user, err
}
