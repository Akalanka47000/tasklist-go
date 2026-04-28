package app

import (
	"go.uber.org/fx"
	"tasklist/modules"
)

// Init provides the fx module for the main fiber app
var Init = append(
	modules.Init,
	fx.Provide(New),
)
