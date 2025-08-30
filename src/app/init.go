package app

import (
	"tasklist/src/modules"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// Init provides the fx module for the main fiber app
var Init = append(
	modules.Init,
	fx.Provide(
		fx.Annotate(New, fx.ResultTags(`name:"app:router"`)),
	),
)

// Params defines the dependencies for the main fiber app
type Params struct {
	fx.In
	Modules *fiber.App `name:"module:router"`
}
