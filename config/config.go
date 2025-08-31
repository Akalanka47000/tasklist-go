// Package config contains the application wise configurations which are loaded from environment variables or a .env file.
package config

type Config struct {
	Port              int    `mapstructure:"PORT"`
	Host              string `mapstructure:"HOST"`
	FrontendBaseURL   string `mapstructure:"FRONTEND_BASE_URL" validate:"url"`
	DatabaseURL       string `mapstructure:"DB_URL" validate:"required,mongodb_connection_string"`
	ServiceRequestKey string `mapstructure:"SERVICE_REQUEST_KEY" validate:"required"` // Key to protect internal routes
	JWTSecret         string `mapstructure:"JWT_SECRET" validate:"required"`
	CI                bool   `mapstructure:"CI"`
	DeploymentEnv     string `mapstructure:"DEPLOYMENT_ENVIRONMENT" validate:"omitempty,oneof=staging production"`
}
