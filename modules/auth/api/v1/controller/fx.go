package controller

import (
	"go.uber.org/fx"
)

// Init provides the fx provider for the auth controller.
var Init = fx.Provide(New)
