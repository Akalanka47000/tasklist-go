package validator

import "go.uber.org/fx"

var Init = fx.Provide(New) // Provide a validator instance to the fx container
