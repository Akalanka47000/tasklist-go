package v1

import (
	"tasklist/src/global"
	"tasklist/src/modules/users/api/v1/dto"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	payload := new(dto.CreateUserReq)
	c.BodyParser(payload)
	res := createUser(c, *payload)
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.CreateUserRes]{
		Message: "User created successfully",
		Data:    res,
	})
}
