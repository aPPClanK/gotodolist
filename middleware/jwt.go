package middleware

import (
	"os"
	"time"

	jwtware "github.com/gofiber/jwt/v4"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

var JWTMiddleware = jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{Key: jwtSecret},
	ContextKey: "user",
})
