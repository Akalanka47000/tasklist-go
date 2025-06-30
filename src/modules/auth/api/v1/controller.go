package v1

import (
	"tasklist/src/global"
	"tasklist/src/modules/auth/api/v1/dto"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	payload := new(dto.LoginRequest)
	c.BodyParser(payload)
	user, accessToken, refreshToken := login(c, payload.Email, payload.Password)
	return c.JSON(global.Response[dto.LoginResponse]{
		Data:    &user,
		Message: "Login successfull!",
	})
}

func RegisterHandler(c *fiber.Ctx) error {
	payload := new(dto.RegisterRequest)
	c.BodyParser(payload)
	res := registerUser(c, *payload)
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.LoginResponse]{
		Data:    res,
		Message: "Registration successfull!",
	})
}
