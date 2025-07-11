package modules

import (
	// "tasklist/src/middleware"
	//"tasklist/src/modules/auth"
	"tasklist/src/modules/users"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	modules := fiber.New()

	//modules.Mount("/", auth.New())

	// modules.All("/", middleware.Protect)

	modules.Mount("/", users.New())

	return modules
}
