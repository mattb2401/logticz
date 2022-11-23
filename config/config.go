package config

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"gopkg.in/yaml.v2"
	"os"
)

//Load Parses []byte data in to default config
func Load(content []byte, config *DefaultConfig) error {
	if err := yaml.Unmarshal(content, config); err != nil {
		return err
	}
	return nil
}

//LoadFile reads file buffer for config file
func LoadFile(filePath string, logger log.Logger) (*DefaultConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		level.Error(logger).Log("msg", "Failed to read config file")
		return nil, err
	}
	cfg := DefaultConfig{}
	if err := Load(data, &cfg); err != nil {
		level.Error(logger).Log("msg", "Failed to unmarshal config yaml")
		return nil, err
	}
	return &cfg, nil
}

//DefaultConfig is configuration for logticz
type DefaultConfig struct {
	ServerConfig struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server_cfg"`
	DatabaseConfig struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Schema   string `yaml:"schema"`
	} `yaml:"database_cfg"`
}
