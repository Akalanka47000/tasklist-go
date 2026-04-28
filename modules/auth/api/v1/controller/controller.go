package controller

import (
	"tasklist/global"
	"tasklist/middleware"
	"tasklist/modules/auth/api/v1/dto"
	authsvc "tasklist/modules/auth/api/v1/service/contracts"
	"tasklist/modules/auth/utils/session"
	. "tasklist/modules/users/api/v1/models"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

type Controller struct {
	service authsvc.Service
}

func New(service authsvc.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// @Summary		Login
// @Description	Authenticate user and return a pair of tokens as cookies
// @Tags			Auth V1
// @Accept			json
// @Produce		json
// @Param			data	body		dto.LoginRequest	true	"Login credentials"
// @Success		200		{object}	global.Response[dto.LoginResponse]
// @Router			/v1/auth/login [post]
func (c *Controller) Login(ctx *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.LoginRequest](ctx)
	user, accessToken, refreshToken := c.service.Login(ctx.Context(), req.Email, req.Password)
	session.SetCookieCredentials(ctx, accessToken, refreshToken)
	return ctx.JSON(global.Response[dto.LoginResponse]{
		Data:    &user,
		Message: "Login successful!",
	})
}

// @Summary		Register
// @Description	Register a new user and return a pair of tokens as cookies
// @Tags			Auth V1
// @Accept			json
// @Produce		json
// @Param			data	body		dto.RegisterRequest	true	"Registration data"
// @Success		201		{object}	global.Response[dto.RegisterResponse]
// @Router			/v1/auth/register [post]
func (c *Controller) Register(ctx *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.RegisterRequest](ctx)
	user, accessToken, refreshToken := c.service.Register(ctx.Context(), *req)
	session.SetCookieCredentials(ctx, accessToken, refreshToken)
	return ctx.Status(fiber.StatusCreated).JSON(global.Response[dto.RegisterResponse]{
		Data:    &user,
		Message: "Registration successful!",
	})
}

// @Summary		Get current user
// @Description	Get the currently authenticated user
// @Tags			Auth V1
// @Accept			json
// @Produce		json
// @Success		200	{object}	global.Response[User]
// @Router			/v1/auth/current [get]
func (c *Controller) CurrentUser(ctx *fiber.Ctx) error {
	user := lo.Cast[*User](ctx.Locals(global.CtxUser))
	return ctx.JSON(global.Response[User]{
		Data:    user,
		Message: "Auth user fetched successfully!",
	})
}

// @Summary		Logout
// @Description	Logout the current user
// @Tags			Auth V1
// @Accept			json
// @Produce		json
// @Success		200	{object}	global.Response[any]
// @Router			/v1/auth/logout [post]
func (c *Controller) Logout(ctx *fiber.Ctx) error {
	session.ClearCookieCredentials(ctx)
	return ctx.Status(fiber.StatusOK).JSON(global.Response[any]{
		Message: "Logout successful!",
	})
}
