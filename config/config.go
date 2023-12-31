package config

import (
	"log"
	"os"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

const (
	ProductionEnv = "production"
	TestEnv       = "testing"

	DatabaseTimeout    = 5 * time.Second
	ProductCachingTime = 10 * time.Minute
)

type Schema struct {
	Environment string `env:"environment"`
	Port        int    `env:"port"`
	AuthSecret  string `env:"auth_secret"`
	DatabaseURI string `env:"database_uri"`
	SmtpHost    string `env:"smtp_host"`
	SmtpUser    string `env:"smtp_user"`
	SmtpPass    string `env:"smtp_pass"`
	SmtpPort    int    `env:"smtp_port"`
	EmailFrom   string `env:"email_from"`
}

var (
	cfg Schema
)

func init() {
	environment := os.Getenv("environment")
	err := godotenv.Load("config/config.yaml")
	if err != nil && environment != TestEnv {
		log.Fatalf("Error on load configuration file, error: %v", err)
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error on parsing configuration file, error: %v", err)
	}
}

func GetConfig() *Schema {
	return &cfg
}
