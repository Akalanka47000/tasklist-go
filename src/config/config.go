package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

type Config struct {
	Port        int    `mapstructure:"PORT"`
	Host        string `mapstructure:"HOST"`
	DatabaseUrl string `mapstructure:"DB_URL" validate:"required"`
	JWTSecret   string `mapstructure:"JWT_SECRET" validate:"required"`
}

var Env *Config

func setDefaults() {
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("HOST", "0.0.0.0")
}

func Load() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&Env); err != nil {
		log.Fatal(err)
	}

	if errs := validator.New().Struct(Env); errs != nil {
		log.Fatal("Invalid environment configuration\n", errs)
	}
}
