package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AutomationEnabled bool `yaml:"automation_enabled"`
	DailyLimit        int  `yaml:"daily_limit"`
}

func Load() Config {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatal("Unable to read config.yml")
	}

	var cfg Config
	yaml.Unmarshal(data, &cfg)
	return cfg
}
