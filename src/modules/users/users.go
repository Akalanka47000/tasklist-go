package users

import (
	"tasklist/src/middleware"
	v1 "tasklist/src/modules/users/api/v1"

	"github.com/akalanka47000/go-modkit/routing"
	"github.com/gofiber/fiber/v2"
)

var versioned = routing.VersionablePrefix("users")

func RegisterRoutes(router fiber.Router) {
	router.All("/*", middleware.Internal)

	router.Route(versioned(1), v1.RegisterRoutes)
}
