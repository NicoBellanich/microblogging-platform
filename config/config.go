package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env string

const (
	Prod  Env = "prod"
	Test  Env = "test"
	Local Env = "local"
)

type Config struct {
	Env Env
}

// IsProdEnv returns if is a production enviroment configuration
func (c *Config) IsProdEnv() bool {
	return c.Env == Prod
}

// IsTestEnv returns if is a test enviroment configuration
func (c *Config) IsTestEnv() bool {
	return c.Env == Test
}

// IsLocalEnv returns if is a localenv configuration
func (c *Config) IsLocalEnv() bool {
	return c.Env == Local
}

// getEnv looks for an env variable or returns default value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func TriggerEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("⚠️  Could't find .env file")
	}
}

// Load load file .env (if exist) and then build the config
func Load() *Config {

	TriggerEnv()

	envStr := getEnv("ENVIRONMENT", "local")

	env := Env(envStr)
	switch env {
	case Local, Test, Prod:
		log.Printf("Environment ['%s'] loaded ok", envStr)
	default:
		log.Printf("Unknown ENVIRONMENT '%s', falling back to 'local'", envStr)
		env = Local
	}

	return &Config{
		Env: env,
	}
}
