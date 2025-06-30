package v1

import (
	m "tasklist/src/middleware"
	"tasklist/src/modules/auth/api/v1/dto"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	v1 := fiber.New()
	v1.Post("/login", m.Zelebrate[dto.LoginRequest](m.ZelebrateSegmentBody), LoginHandler)
	v1.Post("/register", m.Zelebrate[dto.RegisterRequest](m.ZelebrateSegmentBody), RegisterHandler)
	return v1
}
