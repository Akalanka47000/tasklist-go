package session

import (
	"tasklist/config"

	"github.com/gofiber/fiber/v2"
)

const (
	AccessTokenCookieName  = "access_token"
	RefreshTokenCookieName = "refresh_token"
)

// SetCookieCredentials sets the access and refresh tokens as HTTP-only cookies.
func SetCookieCredentials(ctx *fiber.Ctx, accessToken, refreshToken string) {
	ctx.Cookie(&fiber.Cookie{
		Name:     AccessTokenCookieName,
		Value:    accessToken,
		HTTPOnly: true,
		Secure:   !config.IsLocal(),
		SameSite: fiber.CookieSameSiteStrictMode,
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     RefreshTokenCookieName,
		Value:    refreshToken,
		HTTPOnly: true,
		Secure:   !config.IsLocal(),
		SameSite: fiber.CookieSameSiteStrictMode,
		Path:     "/api/v1/auth",
	})
}

// ClearCookieCredentials invalidates the access and refresh token cookies.
func ClearCookieCredentials(ctx *fiber.Ctx) {
	ctx.Cookie(&fiber.Cookie{
		Name:   AccessTokenCookieName,
		Value:  "",
		MaxAge: -1,
	})
	ctx.Cookie(&fiber.Cookie{
		Name:   RefreshTokenCookieName,
		Value:  "",
		MaxAge: -1,
	})
}
