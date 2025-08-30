package v1

import (
	"context"
	"tasklist/src/modules/auth/api/v1/dto"
	"tasklist/src/modules/users/api/v1/models"
	. "tasklist/src/modules/users/api/v1/models"
	userrepo "tasklist/src/modules/users/api/v1/repository"
	"tasklist/src/utils/hash"
	jwtx "tasklist/src/utils/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

func login(ctx context.Context, email, password string) (User, string, string) {
	user := userrepo.GetUserByEmail(ctx, email)
	if user == nil {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials"))
	}
	passwordsMatch := hash.Compare(password, lo.FromPtr(user.Password))
	if !passwordsMatch {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials"))
	}
	accessToken := jwtx.MustGenerateUserToken(*user, false)
	refreshToken := jwtx.MustGenerateUserToken(*user, true)
	return *user, accessToken, refreshToken
}

func registerUser(ctx context.Context, payload dto.RegisterRequest) (User, string, string) {
	user := userrepo.CreateUser(ctx, models.User{
		Name:     &payload.Name,
		Email:    &payload.Email,
		Password: lo.ToPtr(hash.MustString(payload.Password)),
	})
	accessToken := jwtx.MustGenerateUserToken(user, false)
	refreshToken := jwtx.MustGenerateUserToken(user, true)
	return user, accessToken, refreshToken
}
