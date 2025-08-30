package v1

import (
	"tasklist/src/global"
	"tasklist/src/middleware"
	"tasklist/src/modules/users/api/v1/dto"
	"tasklist/src/modules/users/api/v1/models"

	fq "github.com/elcengine/elemental/plugins/filterquery"
	fqm "github.com/elcengine/elemental/plugins/filterquery/middleware"
	"github.com/gofiber/fiber/v2"
)

func CreateUserHandler(c *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.CreateUserRequest](c)
	result := CreateUser(c.Context(), models.User{
		Email:    &req.Email,
		Name:     &req.Name,
		Password: &req.Password,
	})
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.CreateUserResponse]{
		Data:    &result,
		Message: "User created successfully",
	})
}

func GetUsersHandler(c *fiber.Ctx) error {
	result := GetUsers(c.Context(), c.Locals(fqm.CtxKey).(fq.Result))
	return c.JSON(global.Response[dto.GetUsersReponse]{
		Data:    &result,
		Message: "Users fetched successfully!",
	})
}

func GetUserHandler(c *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.GetUserRequest](c)
	result := GetUserByID(c.Context(), req.ID)
	return c.JSON(global.Response[dto.GetUserResponse]{
		Data:    result,
		Message: "User fetched successfully!",
	})
}

func UpdateUserHandler(c *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.UpdateUserRequest](c)
	result := UpdateUserByID(c.Context(), req.ID, models.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	})
	return c.JSON(global.Response[dto.GetUserResponse]{
		Data:    &result,
		Message: "User updated successfully!",
	})
}

func DeleteUserHandler(c *fiber.Ctx) error {
	req := middleware.ZelebrateRequest[dto.DeleteUserRequest](c)
	DeleteUserByID(c.Context(), c.Params(req.ID))
	return c.Status(fiber.StatusOK).JSON(global.Response[any]{
		Data:    nil,
		Message: "User deleted successfully!",
	})
}
