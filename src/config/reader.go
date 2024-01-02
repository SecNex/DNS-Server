package config

import (
	"encoding/json"
	"os"
)

type ConfigReader struct {
	File string
}

type ConfigFile struct {
	Server   ServerConfig   `json:"server"`
	Resolver ResolverConfig `json:"resolver"`
	Backend  BackendConfig  `json:"backend"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ResolverConfig struct {
	Host       string           `json:"host"`
	Port       int              `json:"port"`
	Forwarding ForwardingConfig `json:"forwarding"`
	Blocker    BlockerConfig    `json:"blocker"`
}

type BackendConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type ForwardingConfig struct {
	Enabled bool   `json:"enabled"`
	Server  string `json:"server"`
}

type BlockerConfig struct {
	Lists []BlockListConfig `json:"lists"`
}

type BlockListConfig struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func NewConfigReader(file string) ConfigReader {
	return ConfigReader{
		File: file,
	}
}

func (c *ConfigReader) Read() (ConfigFile, bool) {
	// Check if file exists
	if _, err := os.Stat(c.File); os.IsNotExist(err) {
		return ConfigFile{}, false
	}
	file, err := os.ReadFile(c.File)
	if err != nil {
		return ConfigFile{}, false
	}

	// Parse JSON
	var config ConfigFile
	err = json.Unmarshal(file, &config)
	if err != nil {
		return ConfigFile{}, false
	}

	return config, true
}
