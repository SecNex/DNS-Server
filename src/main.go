package main

import (
	"fmt"
	"os"

	"github.com/secnex/dns-server/config"
)

func main() {
	configFile := "config.json"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}

	// Read config
	file := config.NewConfigReader(configFile)
	config, found := file.Read()
	if !found {
		fmt.Printf("Config file %s not found\n", configFile)
		os.Exit(1)
	}
	registry := config.NewRegistry()
	registry.Start()
}
