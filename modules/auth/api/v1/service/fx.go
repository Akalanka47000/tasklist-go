package service

import (
	"go.uber.org/fx"
)

// Init provides the fx provider for the auth service
var Init = fx.Provide(New)
