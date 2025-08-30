package v1

import (
	m "tasklist/src/middleware"
	"tasklist/src/modules/auth/api/v1/controller"
	"tasklist/src/modules/auth/api/v1/dto"
	"tasklist/src/modules/auth/api/v1/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Init = []fx.Option{
	controller.Init,
	service.Init,
	fx.Provide(
		fx.Annotate(new, fx.ResultTags(`name:"auth:router.v1"`)),
	),
}

type Params struct {
	fx.In
	Controller *controller.Controller
}

// new creates a sub fiber app with auth routes
func new(params Params) *fiber.App {
	app := fiber.New()

	app.Post("/login", m.Zelebrate[dto.LoginRequest](m.ZelebrateSegmentBody), params.Controller.Login)
	app.Post("/register", m.Zelebrate[dto.RegisterRequest](m.ZelebrateSegmentBody), params.Controller.Register)
	app.Get("/current", m.Protect, params.Controller.CurrentUser)
	app.Post("/logout", m.Protect, params.Controller.Logout)

	return app
}
