package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	GCP struct {
		ProjectID    string `toml:"project_id"`
		Zone         string `toml:"zone"`
		InstanceName string `toml:"instance_name"`
		MachineType  string `toml:"machine_type"`
		DiskSizeGB   int    `toml:"disk_size_gb"`
	} `toml:"gcp"`
}

var cfg Config

func Init() error {
	configPaths := []string{
		"./config/config.toml",
		"./config.toml",
		"config.toml",
	}

	var configFile string
	for _, path := range configPaths {
		if _, err := os.Stat(path); err == nil {
			configFile = path
			break
		}
	}

	if configFile == "" {
		return fmt.Errorf("config file not found: please copy config.sample.toml to config.toml and set your values")
	}

	if _, err := toml.DecodeFile(configFile, &cfg); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	return nil
}

func Get() *Config {
	return &cfg
}
