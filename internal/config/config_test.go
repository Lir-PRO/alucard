package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Config struct {
	AppName string `yaml:"app_name"`
}

func TestLoadConfig(t *testing.T) {
	config, _ := LoadConfig[Config]("config_test.yml")
	assert.Equal(t, "Alucard_Deployer", config.AppName)
}
