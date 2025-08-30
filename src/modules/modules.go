package modules

import (
	"tasklist/src/modules/auth"
	"tasklist/src/modules/users"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	modules := fiber.New()

	modules.
		Route("/", auth.RegisterRoutes).
		Route("/", users.RegisterRoutes)

	return modules
}
