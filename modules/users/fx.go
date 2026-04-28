package users

import (
	"go.uber.org/fx"
	v1 "tasklist/modules/users/api/v1"
)

// Init provides the fx module for the user module
var Init = append(
	v1.Init,
	fx.Provide(new),
)
