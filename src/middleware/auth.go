package middleware

import (
	"tasklist/src/config"
	"tasklist/src/global"
	"tasklist/src/modules/auth/utils/session"
	jwtx "tasklist/src/utils/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// Extracts and validates a JWT bearer token from the Authorization header, if present.
// If valid, the user information is stored in the request context.
// If invalid or missing, an error is stored in the context for later handling.
func Sentinel(ctx *fiber.Ctx) error {
	token := ctx.Cookies(session.AccessTokenCookieName)
	if token == "" {
		ctx.Locals(global.CtxAuthorizerError, fiber.NewError(fiber.StatusUnauthorized, "Missing auth token"))
		return ctx.Next()
	}
	user, err := jwtx.ValidateUserToken(token)
	if err != nil {
		ctx.Locals(global.CtxAuthorizerError, err)
		return ctx.Next()
	}
	ctx.Locals(global.CtxUser, user)
	return ctx.Next()
}

// Protects an API route by checking if the request has been recognized by the Sentinel middleware.
func Protect(ctx *fiber.Ctx) error {
	authorizerError := ctx.Locals(global.CtxAuthorizerError)
	if authorizerError != nil {
		panic(authorizerError)
	}
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
