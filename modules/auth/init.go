package auth

import (
	v1 "tasklist/modules/auth/api/v1"

	"github.com/akalanka47000/go-modkit/routing"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var versioned = routing.VersionablePrefix("auth")

// Init provides the fx module for the auth API
var Init = append(
	v1.Init,
	fx.Provide(new),
)

// Params defines the dependencies for the auth module
type Params struct {
	fx.In
	V1 *fiber.App `name:"auth:router.v1"`
}

type Router struct {
	V1 *fiber.App
}

// new creates an auth module router with versioned sub fiber apps
func new(params Params) *Router {
	return &Router{
		V1: params.V1,
	}
}

// ConfigureRoutes registers versioned routes for the entire auth module
func (r *Router) ConfigureRoutes(app *fiber.App) {
	app.Mount(versioned(1), r.V1)
}
