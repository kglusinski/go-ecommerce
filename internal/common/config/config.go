package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type ServiceConfig struct {
	Env       string `envconfig:"ENVIRONMENT" default:"dev"`
	SentryDSN string `envconfig:"SENTRY_DSN"`
}

func (cfg *ServiceConfig) EnvPrefix() string {
	return "ECOMMERCE"
}

func InitConfig() ServiceConfig {
	var cfg ServiceConfig

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load config file: %s", err.Error())
	}

	err = envconfig.Process(cfg.EnvPrefix(), &cfg)
	if err != nil {
		log.Fatalf("could not parse config: %s", err.Error())
	}

	log.Printf("config loaded: %+v", cfg)

	return cfg
}
