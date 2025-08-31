package app

import (
	"tasklist/config"
	"tasklist/global"
	"tasklist/middleware"

	"github.com/gofiber/contrib/swagger"

	elemental "github.com/elcengine/elemental/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var ServiceName = "Todo Service"

// Creates and returns a Fiber application with middleware, routes, and database connection.
func New(params Params) *fiber.App {
	elemental.Connect(config.Env.DatabaseURL)

	app := fiber.New(fiber.Config{
		AppName:      ServiceName,
		ErrorHandler: middleware.ErrorHandler,
		BodyLimit:    512 * 1024, // 512 KB,
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: middleware.StackTraceHandler,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.Env.FrontendBaseUrl,
		AllowCredentials: true,
	}))

	app.Use(helmet.New())

	app.Use(requestid.New(requestid.Config{
		Header:     global.HdrXCorrelationID,
		ContextKey: global.CtxCorrelationID,
	}))

	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(global.Response[*interface{}]{
				Message: "Too many requests, please try again later",
			})
		},
	}))

	app.Use(middleware.Sentinel)

	app.Use(middleware.Zapped)

	app.Use(middleware.Injectors...)

	app.Use(middleware.HealthCheck(middleware.HealthCheckOptions{
		Service: &ServiceName,
		CheckFunctions: map[string]func() bool{
			"database": func() bool {
				return elemental.Ping() == nil
			},
		},
	}))

	if config.IsLocal() {
		app.Use(swagger.New(swagger.Config{
			Path:     "/docs",
			FilePath: "./docs/swagger.json",
			CacheAge: 5, // We want to always serve the latest documentation in local/dev environments
		}))
	}

	app.Mount("/api", params.Modules)

	return app
}
