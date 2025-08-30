package v1

import (
	"github.com/gofiber/fiber/v2"
	"tasklist/src/global"
	"tasklist/src/middleware"
	"tasklist/src/modules/auth/api/v1/dto"
	"tasklist/src/modules/auth/utils/session"
	. "tasklist/src/modules/users/api/v1/models"
)

func LoginHandler(c *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.LoginRequest](c)
	user, accessToken, refreshToken := login(c.Context(), req.Email, req.Password)
	session.SetCookieCredentials(c, accessToken, refreshToken)
	return c.JSON(global.Response[dto.LoginResponse]{
		Data:    &user,
		Message: "Login successfull!",
	})
}

func RegisterHandler(c *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.RegisterRequest](c)
	user, accessToken, refreshToken := registerUser(c.Context(), *req)
	session.SetCookieCredentials(c, accessToken, refreshToken)
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.LoginResponse]{
		Data:    &user,
		Message: "Registration successfull!",
	})
}

func CurrentUserHandler(c *fiber.Ctx) error {
	user := c.Locals(global.CtxUser).(*User)
	return c.JSON(global.Response[User]{
		Data:    user,
		Message: "Auth user fetched successfully!",
	})
}

func LogoutHandler(c *fiber.Ctx) error {
	session.ClearCookieCredentials(c)
	return c.Status(fiber.StatusOK).JSON(global.Response[any]{
		Message: "Logout successfull!",
	})
}
