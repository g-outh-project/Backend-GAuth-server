package utils

import (
	"errors"
	"time"

	"github.com/Backend-GAuth-server/dto"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// test jwtSignKey
var testSignKey = []byte("TestForFasthttpWithJWT")

// Credential type
type userCredential struct {
	Uid               string `json:"uid"`
	Id                string `json:"id"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	School            string `json:"school"`
	Birth             string `json:"birth"`
	Nickname          string `json:"nickname"`
	HashedAccessToken string `json:"hashedAccessToken"`
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
func AccessToken(data dto.JWTSource) string {
	// Generate Token object
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		Id:       data.Id,
		Name:     data.Name,
		Nickname: data.Nickname,
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
func RefreshToken(data dto.JWTSource) string {
	// Generate Token object
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		HashedAccessToken: data.HashedAccessToken,
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
