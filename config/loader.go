package config

import (
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
	"github.com/spf13/viper"
)

var Env *Config // Global variable to hold application wide configuration

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
