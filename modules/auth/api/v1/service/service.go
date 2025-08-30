package service

import (
	"context"
	"tasklist/modules/auth/api/v1/dto"
	contracts "tasklist/modules/auth/api/v1/service/contracts"
	"tasklist/modules/users/api/v1/models"
	users "tasklist/modules/users/api/v1/service/contracts"
	"tasklist/utils/hash"
	jwtx "tasklist/utils/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// service implements the Service interface for authentication logic.
type service struct {
	userService users.Service
}

func new(params Params) contracts.Service {
	return &service{userService: params.UserService}
}

func (s *service) Login(ctx context.Context, email, password string) (models.User, string, string) {
	user := s.userService.GetUserByEmail(ctx, email)
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

func (s *service) Register(ctx context.Context, payload dto.RegisterRequest) (models.User, string, string) {
	user := s.userService.CreateUser(ctx, models.User{
		Name:     &payload.Name,
		Email:    &payload.Email,
		Password: lo.ToPtr(hash.MustString(payload.Password)),
	})
	accessToken := jwtx.MustGenerateUserToken(user, false)
	refreshToken := jwtx.MustGenerateUserToken(user, true)
	return user, accessToken, refreshToken
}
