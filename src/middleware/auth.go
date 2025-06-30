package middleware

import (
	"tasklist/src/config"
	"tasklist/src/global"
	"tasklist/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

func Protect(ctx *fiber.Ctx) error {
	token := ctx.Get(fiber.HeaderAuthorization)
	if token == "" {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Missing bearer token"))
	}
	user := utils.ValidateUserJWTToken(token[len("Bearer "):])
	ctx.Locals("user", user)
	return ctx.Next()
}

// Protects an API route by checking if the request contains a valid service request key.
func Internal(ctx *fiber.Ctx) error {
	if config.Env.ServiceRequestKey == "" {
		return ctx.Next()
	}
	if lo.CoalesceOrEmpty(ctx.Get(global.HdrXServiceRequestKey), ctx.Query("token")) !=
		config.Env.ServiceRequestKey {
		panic(fiber.NewError(fiber.StatusForbidden, "You are not permitted to access this resource"))
	}
	return ctx.Next()
}
