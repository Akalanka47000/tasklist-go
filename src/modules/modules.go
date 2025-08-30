package modules

import (
	//"tasklist/src/modules/auth"
	"tasklist/src/modules/auth"
	"tasklist/src/modules/users"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

// Init provides the fx module for the user API
var Init = append(
	lo.Flatten(
		[][]fx.Option{
			auth.Init,
			users.Init,
		},
	),
	fx.Provide(
		fx.Annotate(new, fx.ResultTags(`name:"module:router"`)),
	),
)

// Params defines the dependencies for the user API
type Params struct {
	fx.In
	Auth  *auth.Router
	Users *users.Router
}

func new(params Params) *fiber.App {
	app := fiber.New()

	params.Auth.ConfigureRoutes(app)

	params.Users.ConfigureRoutes(app)

	return app
}
