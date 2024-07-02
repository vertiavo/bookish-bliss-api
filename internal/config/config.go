package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
)

// Config contains the configuration for the application
type Config struct {
	DBHost     string `env:"DB_HOST,required"`
	DBPort     int    `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`

	Port  int  `env:"PORT" envDefault:"8080"`
	Debug bool `env:"DEBUG" envDefault:"false"`
}

// LoadConfig loads the configuration from the environment
func LoadConfig() *Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	return &cfg
}

// DatabaseURL returns the URL for the database connection
func (cfg *Config) DatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
}
