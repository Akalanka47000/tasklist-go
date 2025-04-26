package users

import (
	"tasklist/src/middleware"
	v1 "tasklist/src/modules/users/api/v1"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	users := fiber.New()
	users.All("/*", middleware.AdminProtect)
	users.Mount("/v1", v1.New())
	return users
}
