package controller

import (
	"tasklist/src/modules/users/api/v1/service"

	"go.uber.org/fx"
)

var Init = fx.Provide(New) // Init provides the fx provider for the user controller

// Params defines the dependencies for the controller
type Params struct {
	fx.In
	Service service.Service
}
