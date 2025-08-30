package auth

import (
	v1 "tasklist/src/modules/auth/api/v1"

	"github.com/akalanka47000/go-modkit/routing"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var versioned = routing.VersionablePrefix("auth")

// Init provides the fx module for the auth API
var Init = append(
	v1.Init,
	fx.Provide(
		fx.Annotate(New, fx.ResultTags(`name:"auth:router"`)),
	),
)

// Params defines the dependencies for the auth API
type Params struct {
	fx.In
	V1 *fiber.App `name:"auth:router.v1"`
}

// New creates a sub fiber app with auth routes
func New(params Params) *fiber.App {
	app := fiber.New()

	app.Mount(versioned(1), params.V1)

	return app
}
