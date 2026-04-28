package v1

import (
	"github.com/gofiber/fiber/v2"
	m "tasklist/middleware"
	"tasklist/modules/auth/api/v1/controller"
	"tasklist/modules/auth/api/v1/dto"
)

// Router defines the current module version's sub fiber app
type Router struct{ *fiber.App }

// New creates a sub fiber app with auth routes
func New(ctrl *controller.Controller) *Router {
	app := fiber.New()

	app.Post("/login", m.Zelebrate[dto.LoginRequest](m.ZelebrateSegmentBody), ctrl.Login)
	app.Post("/register", m.Zelebrate[dto.RegisterRequest](m.ZelebrateSegmentBody), ctrl.Register)
	app.Get("/current", m.Protect, ctrl.CurrentUser)
	app.Post("/logout", m.Protect, ctrl.Logout)

	return &Router{app}
}
