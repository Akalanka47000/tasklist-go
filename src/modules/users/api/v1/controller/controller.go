package controller

import (
	"tasklist/src/global"
	"tasklist/src/middleware"
	"tasklist/src/modules/users/api/v1/dto"
	"tasklist/src/modules/users/api/v1/models"
	serviceContracts "tasklist/src/modules/users/api/v1/service/contracts"

	fq "github.com/elcengine/elemental/plugins/filterquery"
	fqm "github.com/elcengine/elemental/plugins/filterquery/middleware"
	"github.com/gofiber/fiber/v2"
)

// Controller handles HTTP requests for user operations
type Controller struct {
	service serviceContracts.Service
}

// New creates a new instance of the user controller
func New(params Params) *Controller {
	return &Controller{
		service: params.Service,
	}
}

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

func (ctrl *Controller) GetUsers(c *fiber.Ctx) error {
	result := ctrl.service.GetUsers(c.Context(), c.Locals(fqm.CtxKey).(fq.Result))
	return c.JSON(global.Response[dto.GetUsersReponse]{
		Data:    &result,
		Message: "Users fetched successfully!",
	})
}

func (ctrl *Controller) GetUserByID(c *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.GetUserRequest](c)
	result := ctrl.service.GetUserByID(c.Context(), req.ID)
	return c.JSON(global.Response[dto.GetUserResponse]{
		Data:    result,
		Message: "User fetched successfully!",
	})
}

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

func (ctrl *Controller) DeleteUserByID(c *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.DeleteUserRequest](c)
	ctrl.service.DeleteUserByID(c.Context(), c.Params(req.ID))
	return c.Status(fiber.StatusOK).JSON(global.Response[any]{
		Data:    nil,
		Message: "User deleted successfully!",
	})
}
