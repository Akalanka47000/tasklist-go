package users

import (
	"tasklist/src/middleware"
	v1 "tasklist/src/modules/users/api/v1"

	"github.com/akalanka47000/go-modkit/routing"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var versioned = routing.VersionablePrefix("users")

// Init provides the fx module for the user module
var Init = append(
	v1.Init,
	fx.Provide(new),
)

// Params defines the dependencies for the user module
type Params struct {
	fx.In
	V1 *fiber.App `name:"users:router.v1"`
}

type Router struct {
	V1 *fiber.App
}

// new creates a user module router with versioned sub fiber apps
func new(params Params) *Router {
	return &Router{
		V1: params.V1,
	}
}

// Registers versioned routes for the entire user module
func (r *Router) ConfigureRoutes(app *fiber.App) {
	app.All("/*", middleware.Internal)

	app.Mount(versioned(1), r.V1)
}
