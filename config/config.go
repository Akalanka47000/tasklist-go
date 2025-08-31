package config

import (
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/akalanka47000/go-modkit/enums"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
	"github.com/spf13/viper"
)

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

var Env *Config // Global variable to hold application wide configuration

var DeploymentEnv = enums.New(struct {
	enums.String
	Staging    string
	Production string
}{}, enums.Lowercase()) // Enum for deployment environments

// Load reads configuration from environment variables or a .env file or system environment variables and populates the Env variable.
func Load() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	_, b, _, _ := runtime.Caller(0)
	viper.AddConfigPath(filepath.Dir(b) + "/..") // Set like this so it can be loaded by test suites as well

	if err := viper.ReadInConfig(); err != nil {
		typ := reflect.TypeOf(Env).Elem()
		for i := range typ.NumField() {
			viper.BindEnv(typ.Field(i).Tag.Get("mapstructure"))
		}
	}

	lo.Try(func() error {
		setDefaults()
		return nil
	})

	if err := viper.Unmarshal(&Env); err != nil {
		log.Fatal(err)
	}

	if errs := validator.New().Struct(Env); errs != nil {
		log.Fatal("Invalid environment configuration\n", errs)
	}
}

// setDefaults sets default values for configuration parameters.
func setDefaults() {
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("FRONTEND_BASE_URL", "http://localhost:5173")
}

// IsProduction returns true if the application is running in production environment.
func IsProduction() bool {
	return Env.DeploymentEnv == DeploymentEnv.Production
}

// IsLocal returns true if the application is running in a local environment (not staging or production).
func IsLocal() bool {
	return Env.DeploymentEnv != DeploymentEnv.Production && Env.DeploymentEnv != DeploymentEnv.Staging
}
