package tests

import (
	"tasklist/config"
)

// Initializes the configuration for the tests.
func Setup() {
	config.Load()
}
