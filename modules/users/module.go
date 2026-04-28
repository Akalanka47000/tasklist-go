package users

import (
	"tasklist/middleware"
	v1 "tasklist/modules/users/api/v1"

	"github.com/akalanka47000/go-modkit/routing"
	"github.com/gofiber/fiber/v2"
)

var versioned = routing.VersionablePrefix("users")

type Router struct {
	V1 *v1.Router
}

// new creates a user module router with versioned sub fiber apps
func new(v1 *v1.Router) *Router {
	return &Router{
		V1: v1,
	}
}

// ConfigureRoutes registers versioned routes for the entire user module
func (r *Router) ConfigureRoutes(app *fiber.App) {
	app.All("/*", middleware.Internal)

	app.Mount(versioned(1), r.V1.App)
}
