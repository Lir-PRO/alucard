package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AppName     string `yaml:"app_name"`
	Description string `yaml:"description"`
	Author      Author `yaml:"author"`
}

type Author struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return config, nil
}
