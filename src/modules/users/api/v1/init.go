package v1

import (
	m "tasklist/src/middleware"
	sdto "tasklist/src/modules/_shared/dto"
	"tasklist/src/modules/users/api/v1/controller"
	"tasklist/src/modules/users/api/v1/dto"
	"tasklist/src/modules/users/api/v1/repository"
	"tasklist/src/modules/users/api/v1/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// Init provides the fx module for the users v1 API
var Init = []fx.Option{
	repository.Init,
	service.Init,
	controller.Init,
	fx.Provide(
		fx.Annotate(New, fx.ResultTags(`name:"users:router.v1"`)),
	),
}

// Params defines the dependencies for the v1 API
type Params struct {
	fx.In
	Controller *controller.Controller
}

// New creates a sub fiber app with user routes
func New(params Params) *fiber.App {
	app := fiber.New()

	app.Post("/", m.Zelebrate[dto.CreateUserRequest](m.ZelebrateSegmentBody), params.Controller.CreateUser)
	app.Get("/", m.Zelebrate[sdto.PaginatedQuery](m.ZelebrateSegmentQuery),
		m.FilterQuery, params.Controller.GetUsers)
	app.Get("/:id", m.Zelebrate[dto.GetUserRequest](m.ZelebrateSegmentParams), params.Controller.GetUserByID)
	app.Patch("/:id", m.Zelebrate[dto.UpdateUserRequest](m.ZelebrateSegmentParams, m.ZelebrateSegmentBody),
		params.Controller.UpdateUserByID)
	app.Delete("/:id", m.Zelebrate[dto.DeleteUserRequest](m.ZelebrateSegmentParams), params.Controller.DeleteUserByID)

	return app
}
