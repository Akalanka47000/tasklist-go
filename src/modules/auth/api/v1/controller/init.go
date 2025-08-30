package controller

import (
	service "tasklist/src/modules/auth/api/v1/service/contracts"

	"go.uber.org/fx"
)

// Init provides the fx provider for the auth controller.
var Init = fx.Provide(new)

// Params defines the dependencies for the auth controller
type Params struct {
	fx.In
	Service service.Service
}
