package modules

import (
	"github.com/samber/lo"
	"go.uber.org/fx"
	"tasklist/modules/auth"
	"tasklist/modules/users"
)

// Init provides the fx module for the user API
var Init = append(
	lo.Flatten(
		[][]fx.Option{
			auth.Init,
			users.Init,
		},
	),
	fx.Provide(New),
)
