package main

import (
	"context"
	"fmt"
	"tasklist/app"
	"tasklist/config"
	"tasklist/global"
	"tasklist/pkg"

	elemental "github.com/elcengine/elemental/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

// @title			Tasklist API
// @version		1.0
// @description	API documentation for the Tasklist service.
// @contact.name	API Support
// @contact.email	support@example.com
// @host			localhost:8080
// @BasePath		/api
func main() {
	config.Load()

	app := fx.New(
		append(
			lo.Flatten(
				[][]fx.Option{
					app.Init,
					pkg.Init,
				},
			),
			fx.Invoke(registerLifecycle),
		)...,
	)

	app.Run()
}

// Sets up the server lifecycle hooks
func registerLifecycle(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Start the server in a goroutine
			go func() {
				err := app.Listen(fmt.Sprintf(":%d", config.Env.Port))
				if err != nil {
					log.Error("Failed to start server", err)
				}
			}()

			log.Info(fmt.Sprintf("Server starting on port %d", config.Env.Port))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Server shutdown initiated")

			// Shutdown the fiber app
			if err := app.Shutdown(); err != nil {
				log.Error("Error during server shutdown", err)
				return err
			}

			// Disconnect from database
			elemental.Disconnect()

			// Execute shutdown hooks
			global.ExecuteShutdownHooks()

			log.Info("Server shutdown complete")
			return nil
		},
	})
}
