package v1

import (
	"github.com/gofiber/fiber/v2"
	m "tasklist/middleware"
	sdto "tasklist/modules/_shared/dto"
	"tasklist/modules/users/api/v1/controller"
	"tasklist/modules/users/api/v1/dto"
)

// Router defines the current module version's sub fiber app
type Router struct{ *fiber.App }

// New creates a sub fiber app with user routes
func New(ctrl *controller.Controller) *Router {
	app := fiber.New()

	app.Post("/", m.Zelebrate[dto.CreateUserRequest](m.ZelebrateSegmentBody), ctrl.CreateUser)
	app.Get("/", m.Zelebrate[sdto.PaginatedQuery](m.ZelebrateSegmentQuery),
		m.FilterQuery, ctrl.GetUsers)
	app.Get("/:id", m.Zelebrate[dto.GetUserRequest](m.ZelebrateSegmentParams), ctrl.GetUserByID)
	app.Patch("/:id", m.Zelebrate[dto.UpdateUserRequest](m.ZelebrateSegmentParams, m.ZelebrateSegmentBody),
		ctrl.UpdateUserByID)
	app.Delete("/:id", m.Zelebrate[dto.DeleteUserRequest](m.ZelebrateSegmentParams), ctrl.DeleteUserByID)

	return &Router{app}
}
