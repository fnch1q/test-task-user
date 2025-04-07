package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	DatabaseHost     string `env:"DB_HOST"`
	DatabaseDB       string `env:"POSTGRES_DB"`
	DatabasePassword string `env:"POSTGRES_PASSWORD"`
	DatabaseUser     string `env:"POSTGRES_USER"`
	DatabasePort     uint   `env:"DB_PORT"`
	ServerPort       string `env:"SERVER_PORT" env-default:"8080"`
}

func NewLoadConfig() (Config, error) {
	var cfg Config
	cleanenv.ReadConfig(".env", &cfg)
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
