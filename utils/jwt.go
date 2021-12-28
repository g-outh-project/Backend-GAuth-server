package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/Backend-GAuth-server/dto"
	"github.com/Backend-GAuth-server/method"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Credential type
type userCredential struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	School            string `json:"school"`
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
		return nil, errors.New("token cannot found")
	}

	// Return token with type []byte
	return jwt, nil
}

// Generate accessToken
func AccessToken(data dto.JWTSource, c *fiber.Ctx) string {
	cid := GetClientId(c)
	keyPair, _ := method.SelectKeyByCid(cid)
	// Generate Token object
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		Id:                data.Id,
		Name:              data.Name,
		School:            data.School,
		Nickname:          data.Nickname,
		HashedAccessToken: data.HashedAccessToken,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(), // 10 Mins
		},
	})

	// jwt Key
	if cid == "" {
		keyPair.Secret = GetSecretKey()
	}
	jwtSignKey := []byte(keyPair.Secret)
	fmt.Println("--------jwt--------")
	fmt.Println(string(jwtSignKey))

	// Sign token
	access, err := accessToken.SignedString(jwtSignKey)
	HandleErr(err)

	return access
}

// Generate refreshToken
func RefreshToken(data dto.JWTSource, c *fiber.Ctx) string {
	cid := GetClientId(c)
	keyPair, _ := method.SelectKeyByCid(cid)

	// Generate Token object
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		Id:                data.Id,
		Name:              data.Name,
		School:            data.School,
		Nickname:          data.Nickname,
		HashedAccessToken: data.HashedAccessToken,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(336 * time.Hour).Unix(), // 14Days
		},
	})

	// jwt Key
	if cid == "" {
		keyPair.Secret = GetSecretKey()
	}
	jwtSignKey := []byte(Hash(keyPair.Secret))

	// Sign token
	refresh, err := refreshToken.SignedString(jwtSignKey)
	HandleErr(err)

	return refresh
}

// Validate token
func ValidateToken(requestToken string, c *fiber.Ctx) (*jwt.Token, *userCredential, error) {
	// Generate Credential object
	user := &userCredential{}

	cid := GetClientId(c)
	keyPair, _ := method.SelectKeyByCid(cid)

	if cid == "" {
		keyPair.Secret = GetSecretKey()
	}
	// jwt Key
	jwtSignKey := []byte(keyPair.Secret)
	fmt.Println("--------jwt--------")
	fmt.Println(string(jwtSignKey))
	// Parse token and validate
	token, err := jwt.ParseWithClaims(requestToken, user, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})
	HandleErr(err)

	return token, user, err
}
