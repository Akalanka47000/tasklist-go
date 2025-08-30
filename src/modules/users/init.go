package users

import (
	"tasklist/src/middleware"
	v1 "tasklist/src/modules/users/api/v1"

	"github.com/akalanka47000/go-modkit/routing"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var versioned = routing.VersionablePrefix("users")

// Init provides the fx module for the user API
var Init = append(
	v1.Init,
	fx.Provide(
		fx.Annotate(New, fx.ResultTags(`name:"users:router"`)),
	),
)

// Params defines the dependencies for the user API
type Params struct {
	fx.In
	V1 *fiber.App `name:"users:router.v1"`
}

// New creates a sub fiber app with user routes
func New(params Params) *fiber.App {
	app := fiber.New()

	app.All("/*", middleware.Internal)

	app.Mount(versioned(1), params.V1)

	return app
}
