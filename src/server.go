package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"tasklist/src/config"
)

func main() {
	log.Fatal(app().Listen(fmt.Sprintf(":%d", config.Env.Port)))
}
