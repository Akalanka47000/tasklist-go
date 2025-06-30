package users

import (
	"tasklist/src/middleware"
	v1 "tasklist/src/modules/users/api/v1"

	"github.com/akalanka47000/go-modkit/routing"
	"github.com/gofiber/fiber/v2"
)

var versioned = routing.VersionablePrefix("users")

func New() *fiber.App {
	users := fiber.New()
	users.All("/*", middleware.Internal)
	users.Mount(versioned(1), v1.New())
	return users
}
