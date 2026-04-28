package auth

import (
	v1 "tasklist/modules/auth/api/v1"

	"go.uber.org/fx"
)

// Init provides the fx module for the auth API
var Init = append(
	v1.Init,
	fx.Provide(New),
)
