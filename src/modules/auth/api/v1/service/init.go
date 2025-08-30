package service

import (
	users "tasklist/src/modules/users/api/v1/service/contracts"

	"go.uber.org/fx"
)

// Init provides the fx provider for the auth service
var Init = fx.Provide(New)

// Params defines the dependencies for the auth service
type Params struct {
	fx.In
	UserService users.Service
}
