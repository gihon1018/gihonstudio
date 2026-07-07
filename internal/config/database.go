package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type DBConfig struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	DBName          string `yaml:"db_name"`
	SSLMode         string `yaml:"ssl_mode"`
	TimeZone        string `yaml:"timezone"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

type AppConfig struct {
	Database DBConfig `yaml:"database"`
}

func LoadConfig() (*AppConfig, error) {
	return LoadConfigFromPath("config/database.yaml")
}

func LoadConfigFromPath(path string) (*AppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg AppConfig

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
