package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Env  string     `ENV:"ENV" env-default:"local"`
	GRPC GRPCConfig `ENV:"grpc"`
}

type GRPCConfig struct {
	Port int `ENV:"PORT"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("config path is empty: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	godotenv.Load("../../domain/config/.env")

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
