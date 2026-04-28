package v1

import (
	"go.uber.org/fx"
	"tasklist/modules/auth/api/v1/controller"
	"tasklist/modules/auth/api/v1/service"
)

var Init = []fx.Option{
	controller.Init,
	service.Init,
	fx.Provide(New),
}
