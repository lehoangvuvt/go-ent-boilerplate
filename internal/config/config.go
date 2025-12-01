package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var _validate = validator.New(validator.WithRequiredStructEnabled())

type Config struct {
	App      AppConfig      `validate:"required"`
	DB       DBConfig       `validate:"required"`
	JWT      JWTConfig      `validate:"required"`
	Redis    RedisConfig    `validate:"required"`
	Mail     MailConfig     `validate:"required"`
	Resend   ResendConfig   `validate:"required"`
	RabbitMQ RabbitMQConfig `validate:"required"`
}

type AppConfig struct {
	Port int `envconfig:"PORT" validate:"required,gt=0" default:"8080"`
}

type DBConfig struct {
	AutoMigrate bool   `envconfig:"AUTO_MIGRATE" default:"true"`
	Host        string `envconfig:"HOST" validate:"required"`
	Port        int    `envconfig:"PORT" validate:"required,gt=0"`
	Name        string `envconfig:"NAME" validate:"required"`
	User        string `envconfig:"USER" validate:"required"`
	Password    string `envconfig:"PASSWORD" validate:"required"`
}

type JWTConfig struct {
	Secret   string `envconfig:"SECRET" validate:"required"`
	Duration int    `envconfig:"DURATION" validate:"required,gt=0"`
}

type RedisConfig struct {
	Address  string `envconfig:"ADDRESS" validate:"required"`
	Password string `envconfig:"PASSWORD"`
}

type RabbitMQConfig struct {
	DSN string `envconfig:"DSN" validate:"required"`
}

type MailConfig struct {
	Host string `envconfig:"HOST" validate:"required"`
	Port int    `envconfig:"PORT" validate:"required,gt=0"`
	User string `envconfig:"USER" validate:"required"`
	Pass string `envconfig:"PASS" validate:"required"`
}

type ResendConfig struct {
	ApiKey string `envconfig:"API_KEY" validate:"required"`
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	var cfg Config

	if err := envconfig.Process("APP", &cfg.App); err != nil {
		return nil, fmt.Errorf("load APP config: %w", err)
	}
	if err := envconfig.Process("DB", &cfg.DB); err != nil {
		return nil, fmt.Errorf("load DB config: %w", err)
	}
	if err := envconfig.Process("JWT", &cfg.JWT); err != nil {
		return nil, fmt.Errorf("load JWT config: %w", err)
	}
	if err := envconfig.Process("REDIS", &cfg.Redis); err != nil {
		return nil, fmt.Errorf("load REDIS config: %w", err)
	}
	if err := envconfig.Process("SMTP", &cfg.Mail); err != nil {
		return nil, fmt.Errorf("load SMTP config: %w", err)
	}
	if err := envconfig.Process("RESEND", &cfg.Resend); err != nil {
		return nil, fmt.Errorf("load RESEND config: %w", err)
	}
	if err := envconfig.Process("RABBITMQ", &cfg.RabbitMQ); err != nil {
		return nil, fmt.Errorf("load RABBITMQ config: %w", err)
	}

	if err := _validate.Struct(&cfg); err != nil {
		return nil, fmt.Errorf("config validation: %w", err)
	}

	return &cfg, nil
}
