package v1

import (
	m "tasklist/src/middleware"
	"tasklist/src/modules/auth/api/v1/dto"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	router.Post("/login", m.Zelebrate[dto.LoginRequest](m.ZelebrateSegmentBody), LoginHandler)
	router.Post("/register", m.Zelebrate[dto.RegisterRequest](m.ZelebrateSegmentBody), RegisterHandler)
	router.Get("/current", m.Protect, CurrentUserHandler)
	router.Post("/logout", m.Protect, LogoutHandler)
}
