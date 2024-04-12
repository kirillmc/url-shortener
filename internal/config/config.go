package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	configPath = "CONFIG_PATH"
	envPath    = "github.com/kirillmc/url-shortener/.env"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"local"`
	StoragePath string `yaml:"storage-path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

// Must - фугкция вместо возврата ошибки - паникует
func MustLoad() *Config {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("failed to get .env: %v", err)
	}

	configPath := os.Getenv(configPath)
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if dile exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
