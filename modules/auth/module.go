package auth

import (
	v1 "tasklist/modules/auth/api/v1"

	"github.com/akalanka47000/go-modkit/routing"
	"github.com/gofiber/fiber/v2"
)

var versioned = routing.VersionablePrefix("auth")

type Router struct {
	V1 *v1.Router
}

// New creates an auth module router with versioned sub fiber apps
func New(v1 *v1.Router) *Router {
	return &Router{
		V1: v1,
	}
}

// ConfigureRoutes registers versioned routes for the entire auth module
func (r *Router) ConfigureRoutes(app *fiber.App) {
	app.Mount(versioned(1), r.V1.App)
}
