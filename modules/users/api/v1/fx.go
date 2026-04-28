package v1

import (
	"go.uber.org/fx"
	"tasklist/modules/users/api/v1/controller"
	"tasklist/modules/users/api/v1/repository"
	"tasklist/modules/users/api/v1/service"
)

// Init provides the fx module for the users v1 API
var Init = []fx.Option{
	repository.Init,
	service.Init,
	controller.Init,
	fx.Provide(New),
}
