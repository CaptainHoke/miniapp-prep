package config

import "time"

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
