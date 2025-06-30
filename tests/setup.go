package tests

import (
	"tasklist/src/config"
)

// Initializes the configuration for the tests.
func Setup() {
	config.Load()
}
