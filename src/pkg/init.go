package pkg

import (
	"tasklist/src/pkg/validator"

	"go.uber.org/fx"
)

// Init provides the fx module for all package extensions
var Init = []fx.Option{
	validator.Init,
}
