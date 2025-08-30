// Package ts provides setup and teardown functions for the test suite.
package ts

import (
	"tasklist/app"
	"tasklist/config"
	"tasklist/pkg"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"github.com/smartystreets/goconvey/convey"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func Prepare(t *testing.T, c convey.C, options ...fx.Option) (fiberApp *fiber.App) {
	t.Helper()

	config.Load()

	app := fxtest.New(t,
		lo.Flatten(
			[][]fx.Option{
				{
					fx.StartTimeout(time.Minute * 5),
					fx.StopTimeout(time.Minute * 5),
					fx.Populate(
						&fiberApp,
					),
				},
				app.Init,
				pkg.Init,
				options,
			},
		)...,
	).RequireStart()

	c.Reset(app.RequireStop)

	return
}
