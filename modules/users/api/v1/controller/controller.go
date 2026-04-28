package controller

import (
	"tasklist/global"
	"tasklist/middleware"
	"tasklist/modules/users/api/v1/dto"
	"tasklist/modules/users/api/v1/models"
	usersvc "tasklist/modules/users/api/v1/service/contracts"

	fq "github.com/elcengine/elemental/plugins/filterquery"
	fqm "github.com/elcengine/elemental/plugins/filterquery/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

type Controller struct {
	service usersvc.Service
}

func new(service usersvc.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// @Summary		Create a new user
// @Description	Create a new user with the input payload
// @Tags			Users V1
// @Accept			json
// @Produce		json
// @Param			data	body		dto.CreateUserRequest	true	"User data"
// @Success		201		{object}	global.Response[dto.CreateUserResponse]
// @Router			/v1/users [post]
func (c *Controller) CreateUser(ctx *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.CreateUserRequest](ctx)
	result := c.service.CreateUser(ctx.Context(), models.User{
		Email:    &req.Email,
		Name:     &req.Name,
		Password: &req.Password,
	})
	return ctx.Status(fiber.StatusCreated).JSON(global.Response[dto.CreateUserResponse]{
		Data:    &result,
		Message: "User created successfully",
	})
}

// @Summary		Get all users
// @Description	Get a list of all users
// @Tags			Users V1
// @Accept			json
// @Produce		json
// @Param			page	query		int	false	"Page number"
// @Param			limit	query		int	false	"Page size"
// @Success		200		{object}	global.Response[dto.GetUsersReponse]
// @Router			/v1/users [get]
func (c *Controller) GetUsers(ctx *fiber.Ctx) error {
	result := c.service.GetUsers(ctx.Context(), lo.Cast[fq.Result](ctx.Locals(fqm.CtxKey)))
	return ctx.JSON(global.Response[dto.GetUsersReponse]{
		Data:    &result,
		Message: "Users fetched successfully!",
	})
}

// @Summary		Get user by ID
// @Description	Get a user by their ID
// @Tags			Users V1
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"User ID"
// @Success		200	{object}	global.Response[dto.GetUserResponse]
// @Router			/v1/users/{id} [get]
func (c *Controller) GetUserByID(ctx *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.GetUserRequest](ctx)
	result := c.service.GetUserByID(ctx.Context(), req.ID)
	return ctx.JSON(global.Response[dto.GetUserResponse]{
		Data:    result,
		Message: "User fetched successfully!",
	})
}

// @Summary		Update user by ID
// @Description	Update a user's information by their ID
// @Tags			Users V1
// @Accept			json
// @Produce		json
// @Param			id		path		string					true	"User ID"
// @Param			data	body		dto.UpdateUserRequest	true	"User data"
// @Success		200		{object}	global.Response[dto.GetUserResponse]
// @Router			/v1/users/{id} [patch]
func (c *Controller) UpdateUserByID(ctx *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.UpdateUserRequest](ctx)
	result := c.service.UpdateUserByID(ctx.Context(), req.ID, models.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	})
	return ctx.JSON(global.Response[dto.GetUserResponse]{
		Data:    &result,
		Message: "User updated successfully!",
	})
}

// @Summary		Delete user by ID
// @Description	Delete a user by their ID
// @Tags			Users V1
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"User ID"
// @Success		200	{object}	global.Response[any]
// @Router			/v1/users/{id} [delete]
func (c *Controller) DeleteUserByID(ctx *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.DeleteUserRequest](ctx)
	c.service.DeleteUserByID(ctx.Context(), ctx.Params(req.ID))
	return ctx.Status(fiber.StatusOK).JSON(global.Response[any]{
		Data:    nil,
		Message: "User deleted successfully!",
	})
}
