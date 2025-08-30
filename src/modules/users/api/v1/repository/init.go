package repository

import (
	"go.uber.org/fx"
)

var Init = fx.Provide(new) // Init provides the fx provider for the user repository
