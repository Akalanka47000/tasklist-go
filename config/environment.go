package config

import (
	"github.com/akalanka47000/go-modkit/enums"
)

var DeploymentEnv = enums.New(struct {
	enums.String
	Staging    string
	Production string
}{}, enums.Lowercase()) // Enum for deployment environments

// IsProduction returns true if the application is running in the production environment.
func IsProduction() bool {
	return Env.DeploymentEnv == DeploymentEnv.Production
}

// IsLocal returns true if the application is running in a local environment (not staging or production).
func IsLocal() bool {
	return Env.DeploymentEnv != DeploymentEnv.Production && Env.DeploymentEnv != DeploymentEnv.Staging
}
