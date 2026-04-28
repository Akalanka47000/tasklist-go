package middleware

import (
	"os"
	"tasklist/global"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// Intercepts the response and populates it with additional fields.
func responseHeaderInjector(ctx *fiber.Ctx) error {
	ctx.Append(global.HdrXHostname, lo.Ok(os.Hostname())) // Useful for debugging
	return ctx.Next()
}

var Injectors = []any{responseHeaderInjector}
