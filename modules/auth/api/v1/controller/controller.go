package controller

import (
	"tasklist/global"
	"tasklist/middleware"
	"tasklist/modules/auth/api/v1/dto"
	service "tasklist/modules/auth/api/v1/service/contracts"
	"tasklist/modules/auth/utils/session"
	. "tasklist/modules/users/api/v1/models"

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

// @Summary		Login
// @Description	Authenticate user and return a pair of tokens as cookies
// @Tags			Auth V1
// @Accept			json
// @Produce		json
// @Param			data	body		dto.LoginRequest	true	"Login credentials"
// @Success		200		{object}	global.Response[dto.LoginResponse]
// @Router			/v1/auth/login [post]
func (ctrl *Controller) Login(c *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.LoginRequest](c)
	user, accessToken, refreshToken := ctrl.service.Login(c.Context(), req.Email, req.Password)
	session.SetCookieCredentials(c, accessToken, refreshToken)
	return c.JSON(global.Response[dto.LoginResponse]{
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
func (ctrl *Controller) Register(c *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.RegisterRequest](c)
	user, accessToken, refreshToken := ctrl.service.Register(c.Context(), *req)
	session.SetCookieCredentials(c, accessToken, refreshToken)
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.RegisterResponse]{
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
func (ctrl *Controller) CurrentUser(c *fiber.Ctx) error {
	user := c.Locals(global.CtxUser).(*User)
	return c.JSON(global.Response[User]{
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
func (ctrl *Controller) Logout(c *fiber.Ctx) error {
	session.ClearCookieCredentials(c)
	return c.Status(fiber.StatusOK).JSON(global.Response[any]{
		Message: "Logout successful!",
	})
}
