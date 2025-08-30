package auth

import (
	v1 "tasklist/src/modules/auth/api/v1"

	"github.com/akalanka47000/go-modkit/routing"
	"github.com/gofiber/fiber/v2"
)

var versioned = routing.VersionablePrefix("auth")

func RegisterRoutes(router fiber.Router) {
	router.Route(versioned(1), v1.RegisterRoutes)
}
