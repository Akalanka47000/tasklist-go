package service

import (
	"tasklist/src/modules/users/api/v1/repository"

	"go.uber.org/fx"
)

var Init = fx.Provide(new) // Init provides the fx provider for the user service

// Params defines the dependencies for the service
type Params struct {
	fx.In
	Repository repository.Repository
}
