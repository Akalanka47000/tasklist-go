package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"tasklist/src/config"
	"tasklist/src/database"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	config.Load()

	app := bootstrapApp()

	go func() {
		err := app.Listen(fmt.Sprintf(":%d", config.Env.Port))
		if err != nil {
			log.Error("Failed to start server", err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c // Waits for termination signal before proceeding

	log.Info("Received SIGTERM. Server shutdown initiated")

	app.Shutdown()

	log.Info("Server shutdown complete. Exiting after 30 seconds")

	time.Sleep(30 * time.Second) // Wait for 30 seconds before closing any open connections

	database.Disconnect()
}
