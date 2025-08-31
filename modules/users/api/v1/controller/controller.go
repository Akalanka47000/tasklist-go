package controller

import (
	"tasklist/global"
	"tasklist/middleware"
	"tasklist/modules/users/api/v1/dto"
	"tasklist/modules/users/api/v1/models"
	serviceContracts "tasklist/modules/users/api/v1/service/contracts"

	fq "github.com/elcengine/elemental/plugins/filterquery"
	fqm "github.com/elcengine/elemental/plugins/filterquery/middleware"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service serviceContracts.Service
}

func new(params Params) *Controller {
	return &Controller{
		service: params.Service,
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param data body dto.CreateUserRequest true "User data"
// @Success 201 {object} global.Response[dto.CreateUserResponse]
// @Router /users [post]
func (ctrl *Controller) CreateUser(c *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.CreateUserRequest](c)
	result := ctrl.service.CreateUser(c.Context(), models.User{
		Email:    &req.Email,
		Name:     &req.Name,
		Password: &req.Password,
	})
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.CreateUserResponse]{
		Data:    &result,
		Message: "User created successfully",
	})
}

// GetUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Page size"
// @Success 200 {object} global.Response[dto.GetUsersReponse]
// @Router /users [get]
func (ctrl *Controller) GetUsers(c *fiber.Ctx) error {
	result := ctrl.service.GetUsers(c.Context(), c.Locals(fqm.CtxKey).(fq.Result))
	return c.JSON(global.Response[dto.GetUsersReponse]{
		Data:    &result,
		Message: "Users fetched successfully!",
	})
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} global.Response[dto.GetUserResponse]
// @Router /users/{id} [get]
func (ctrl *Controller) GetUserByID(c *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.GetUserRequest](c)
	result := ctrl.service.GetUserByID(c.Context(), req.ID)
	return c.JSON(global.Response[dto.GetUserResponse]{
		Data:    result,
		Message: "User fetched successfully!",
	})
}

// UpdateUserByID godoc
// @Summary Update user by ID
// @Description Update a user's information by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param data body dto.UpdateUserRequest true "User data"
// @Success 200 {object} global.Response[dto.GetUserResponse]
// @Router /users/{id} [patch]
func (ctrl *Controller) UpdateUserByID(c *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.UpdateUserRequest](c)
	result := ctrl.service.UpdateUserByID(c.Context(), req.ID, models.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	})
	return c.JSON(global.Response[dto.GetUserResponse]{
		Data:    &result,
		Message: "User updated successfully!",
	})
}

// DeleteUserByID godoc
// @Summary Delete user by ID
// @Description Delete a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} global.Response[any]
// @Router /users/{id} [delete]
func (ctrl *Controller) DeleteUserByID(c *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.DeleteUserRequest](c)
	ctrl.service.DeleteUserByID(c.Context(), c.Params(req.ID))
	return c.Status(fiber.StatusOK).JSON(global.Response[any]{
		Data:    nil,
		Message: "User deleted successfully!",
	})
}
