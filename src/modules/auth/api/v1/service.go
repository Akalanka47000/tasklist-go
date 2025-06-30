package v1

import (
	"fmt"
	"tasklist/src/modules/auth/api/v1/dto"
	user "tasklist/src/modules/users/api/v1"
	"tasklist/src/modules/users/api/v1/models"
	. "tasklist/src/modules/users/api/v1/models"
	userrepo "tasklist/src/modules/users/api/v1/repository"
	"tasklist/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
)

func login(c *fiber.Ctx, email, password string) (User, string, string) {
	user := userrepo.GetUserByEmail(email)
	if user == nil {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials"))
	}
	passwordsMatch := utils.CompareStrHash(password, lo.FromPtr(user.Password))
	if !passwordsMatch {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials"))
	}
	accessToken := utils.GenerateUserJWTToken(*user, false)
	refreshToken := utils.GenerateUserJWTToken(*user, true)
	return *user, accessToken, refreshToken
}

func registerUser(c *fiber.Ctx, payload dto.RegisterRequest) *dto.LoginResponse {
	insertedID := user.Repository().Create(models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: utils.HashStr(payload.Password),
	})
	accessToken := utils.GenerateUserJWTToken(user, false)
	refreshToken := utils.GenerateUserJWTToken(user, true)
	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}
}
