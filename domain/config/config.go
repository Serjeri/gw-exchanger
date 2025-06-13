package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Database   `yaml:"database"`
	HTTPServer `yaml:"http_server"`
}

type Database struct {
	Dburl string `yaml:"dburl"`
}

type HTTPServer struct {
	Addressgrpc string `yaml:"addressgrpc"`
}

func MustLoad() *Config {
	configPath := "../config/config.yml"
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		log.Fatalf("cannot parse YAML config: %s", err)
	}

	return &cfg
}
