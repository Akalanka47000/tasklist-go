package controller

import (
	service "tasklist/src/modules/users/api/v1/service/contracts"

	"go.uber.org/fx"
)

var Init = fx.Provide(new) // Init provides the fx provider for the user controller

// Params defines the dependencies for the controller
type Params struct {
	fx.In
	Service service.Service
}
