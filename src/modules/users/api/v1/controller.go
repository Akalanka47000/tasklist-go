package v1

import (
	"tasklist/src/global"
	"tasklist/src/modules/users/api/v1/dto"
	"tasklist/src/modules/users/api/v1/models"

	fq "github.com/elcengine/elemental/plugins/filterquery"
	fqm "github.com/elcengine/elemental/plugins/filterquery/middleware"
	"github.com/gofiber/fiber/v2"
)

func CreateUserHandler(c *fiber.Ctx) error {
	payload := new(dto.CreateUserRequest)
	c.BodyParser(payload)
	result := CreateUser(models.User{
		Email:    &payload.Email,
		Name:     &payload.Name,
		Password: &payload.Password,
	})
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.CreateUserResponse]{
		Data:    &result,
		Message: "User created successfully",
	})
}

func GetUsersHandler(c *fiber.Ctx) error {
	result := GetUsers(c.Locals(fqm.CtxKey).(fq.Result))
	return c.JSON(global.Response[dto.GetUsersReponse]{
		Data:    &result,
		Message: "Users fetched successfully!",
	})
}

func GetUserHandler(c *fiber.Ctx) error {
	result := GetUserByID(c.Params("id"))
	return c.JSON(global.Response[dto.GetUserResponse]{
		Data:    result,
		Message: "User fetched successfully!",
	})
}

func UpdateUserHandler(c *fiber.Ctx) error {
	payload := new(dto.UpdateUserRequest)
	c.BodyParser(payload)
	result := UpdateUserByID(c.Params("id"), models.User{
		Email:    payload.Email,
		Name:     payload.Name,
		Password: payload.Password,
	})
	return c.JSON(global.Response[dto.GetUserResponse]{
		Data:    &result,
		Message: "User updated successfully!",
	})
}

func DeleteUserHandler(c *fiber.Ctx) error {
	DeleteUserByID(c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(global.Response[any]{
		Data:    nil,
		Message: "User deleted successfully!",
	})
}
