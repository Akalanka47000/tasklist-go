// Custom utility functions for generating and validating JWT tokens.
package jwtx

import (
	"encoding/json"
	"tasklist/src/config"
	. "tasklist/src/modules/users/api/v1/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// MustGenerateUserToken generates a JWT token for the given user.
// If refresh is true, it generates a refresh token with a longer expiry.
func MustGenerateUserToken(user User, refresh bool) string {
	expiry := time.Hour * 1
	if refresh {
		expiry = time.Hour * 24
	}
	claims := jwt.MapClaims{
		"iss":  "Tasklist",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(expiry).Unix(),
		"data": user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Env.JWTSecret))
	if err != nil {
		panic(fiber.NewError(fiber.StatusInternalServerError, "Error generating jwt token"))
	}
	return t
}

// ValidateUserToken validates a given JWT token and parses the user information from it.
func ValidateUserToken(token string) (*User, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(config.Env.JWTSecret), nil
	})
	if err != nil || !parsedToken.Valid {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}
	user := User{}
	jsonString, _ := json.Marshal(claims["data"])
	json.Unmarshal(jsonString, &user)
	return &user, nil
}
