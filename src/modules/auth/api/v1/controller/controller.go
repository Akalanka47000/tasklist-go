package controller

import (
	"tasklist/src/global"
	"tasklist/src/middleware"
	"tasklist/src/modules/auth/api/v1/dto"
	"tasklist/src/modules/auth/api/v1/service"
	"tasklist/src/modules/auth/utils/session"
	. "tasklist/src/modules/users/api/v1/models"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service service.Service
}

func new(params Params) *Controller {
	return &Controller{
		service: params.Service,
	}
}

func (ctrl *Controller) Login(c *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.LoginRequest](c)
	user, accessToken, refreshToken := ctrl.service.Login(c.Context(), req.Email, req.Password)
	session.SetCookieCredentials(c, accessToken, refreshToken)
	return c.JSON(global.Response[dto.LoginResponse]{
		Data:    &user,
		Message: "Login successful!",
	})
}

func (ctrl *Controller) Register(c *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.RegisterRequest](c)
	user, accessToken, refreshToken := ctrl.service.RegisterUser(c.Context(), *req)
	session.SetCookieCredentials(c, accessToken, refreshToken)
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.RegisterResponse]{
		Data:    &user,
		Message: "Registration successful!",
	})
}

func (ctrl *Controller) CurrentUser(c *fiber.Ctx) error {
	user := c.Locals(global.CtxUser).(*User)
	return c.JSON(global.Response[User]{
		Data:    user,
		Message: "Auth user fetched successfully!",
	})
}

func (ctrl *Controller) Logout(c *fiber.Ctx) error {
	session.ClearCookieCredentials(c)
	return c.Status(fiber.StatusOK).JSON(global.Response[any]{
		Message: "Logout successful!",
	})
}
