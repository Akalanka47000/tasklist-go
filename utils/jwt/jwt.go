// Package jwtx provides custom utility functions for generating and validating JWT tokens.
package jwtx

import (
	"encoding/json"
	"tasklist/config"
	. "tasklist/modules/users/api/v1/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/lo"
)

// MustGenerateUserToken generates a JWT token for the given user.
// If refresh is true, it generates a refresh token with a longer expiry.
func MustGenerateUserToken(user User, refresh bool) string {
	expiry := time.Hour * 1
	if refresh {
		expiry = time.Hour * 24
	}
	user.Password = nil // Remove password before embedding in token
	claims := jwt.MapClaims{
		"iss":  "Tasklist",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(expiry).Unix(),
		"data": user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return lo.Must(token.SignedString([]byte(config.Env.JWTSecret)))
}

// ValidateUserToken validates a given JWT token and parses the user information from it.
func ValidateUserToken(token string) (*User, error) {
	invalidTokenErr := fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(config.Env.JWTSecret), nil
	})
	if err != nil || !parsedToken.Valid {
		return nil, invalidTokenErr
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, invalidTokenErr
	}
	user := User{}
	jsonString, err := json.Marshal(claims["data"])
	if err != nil {
		return nil, invalidTokenErr
	}
	json.Unmarshal(jsonString, &user)
	return &user, nil
}
