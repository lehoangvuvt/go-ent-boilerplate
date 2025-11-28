package config

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	AppConfig   AppConfig   `mapstructure:"app" validate:"required"`
	DBConfig    DBConfig    `mapstructure:"db" validate:"required"`
	JWTConfig   JWTConfig   `mapstructure:"jwt" validate:"required"`
	RedisConfig RedisConfig `mapstructure:"redis" validate:"required"`
}

type AppConfig struct {
	Port int `mapstructure:"port" validate:"required"`
}

var _validate = validator.New(validator.WithRequiredStructEnabled())

type DBConfig struct {
	AutoMigrate bool   `mapstructure:"auto_migrate"`
	Host        string `mapstructure:"host" validate:"required"`
	Port        int    `mapstructure:"port" validate:"required,gt=0"`
	Name        string `mapstructure:"name" validate:"required"`
	User        string `mapstructure:"user" validate:"required"`
	Password    string `mapstructure:"password" validate:"required"`
}

type JWTConfig struct {
	Secret   string `mapstructure:"secret" validate:"required"`
	Duration int    `mapstructure:"duration" validate:"required,gt=0"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address" validate:"required"`
	Password string `mapstructure:"password"`
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.SetDefault("app.port", 0)

	v.SetDefault("db.host", "")
	v.SetDefault("db.port", 0)
	v.SetDefault("db.name", "")
	v.SetDefault("db.user", "")
	v.SetDefault("db.password", "")
	v.SetDefault("db.auto_migrate", true)

	v.SetDefault("jwt.secret", "")
	v.SetDefault("jwt.duration", 0)

	v.SetDefault("redis.address", "")
	v.SetDefault("redis.password", "")

	cfg := new(Config)
	if err := v.Unmarshal(cfg); err != nil {
		return nil, err
	}

	if err := validate(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func validate(cfg *Config) error {
	if err := _validate.Struct(cfg); err != nil {
		return fmt.Errorf("config validation: %w", err)
	}
	return nil
}
