package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	config, _ := LoadConfig("config_test.yml")
	assert.Equal(t, "Alucard_Deployer", config.AppName)
}
