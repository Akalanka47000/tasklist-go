package modules

import (
	"tasklist/modules/auth"
	"tasklist/modules/users"

	"github.com/gofiber/fiber/v2"
)

// Router defines the root fiber app that will be used to register all module routes. This is the main entry point for the app.
type Router struct{ *fiber.App }

// New initializes the main module level fiber app and registers all sub module routes.
func New(auth *auth.Router, users *users.Router) *Router {
	app := fiber.New()

	auth.ConfigureRoutes(app)

	users.ConfigureRoutes(app)

	return &Router{app}
}
