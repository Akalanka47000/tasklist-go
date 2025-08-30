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
		fx.Annotate(New, fx.ResultTags(`name:"module:router"`)),
	),
)

// Params defines the dependencies for the user API
type Params struct {
	fx.In
	Auth  *fiber.App `name:"auth:router"`
	Users *fiber.App `name:"users:router"`
}

func New(params Params) *fiber.App {
	modules := fiber.New()

	modules.
		Mount("/", params.Auth).
		Mount("/", params.Users)

	return modules
}
