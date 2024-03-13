package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env:"NK_ENV" env-default:"local" env-required:"true"`
	StoragePath string `yaml:"storage_path" env:"NK_STORAGE_PATH" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:3239"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"15s"`
}

const cfgEnvVar = "NK_CONFIG_PATH"

func MustLoadCfg() Config {
	// TODO: cmd line param
	cfgPath := os.Getenv(cfgEnvVar)

	if cfgPath == "" {
		log.Fatalf("%s not set", cfgEnvVar)
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalf("Config not found: %s", cfgPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		log.Fatalf("Config read error: %s", err)
	}

	return cfg
}
