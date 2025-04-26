package v1

import (
	m "tasklist/src/middleware"
	"tasklist/src/modules/auth/api/v1/dto"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	v1 := fiber.New()
	v1.Post("/login", m.Validate[dto.LoginReq](m.Body), Login)
	v1.Post("/register", m.Validate[dto.RegisterReq](m.Body), Register)
	return v1
}
