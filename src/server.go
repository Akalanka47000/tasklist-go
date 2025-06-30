package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"tasklist/src/app"
	"tasklist/src/config"
	"tasklist/src/global"
	"time"

	"github.com/elcengine/elemental/core"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	config.Load()

	app := app.New()

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

	log.Info("Server shutdown complete. Exiting after 10 seconds")

	time.Sleep(10 * time.Second)

	elemental.Disconnect()

	global.ExecuteShutdownHooks()
}
