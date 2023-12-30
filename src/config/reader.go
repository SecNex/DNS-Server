package config

import (
	"encoding/json"
	"os"

	"github.com/secnex/dns-server/resolver"
)

type ConfigReader struct {
	File string
}

type ConfigFile struct {
	Server   ServerConfig   `json:"server"`
	Resolver ResolverConfig `json:"resolver"`
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

func (c *ConfigFile) NewRegistry() resolver.DnsRegistry {
	reg := resolver.NewRegistry(c.Resolver.Host, c.Resolver.Port)
	reg.Forwarding.SetForwarding(c.Resolver.Forwarding.Enabled, c.Resolver.Forwarding.Server)
	for _, list := range c.Resolver.Blocker.Lists {
		reg.Blocker.AddList(list.Name, list.Url)
	}
	reg.Blocker.Sync()
	return reg
}
