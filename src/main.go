package main

import (
	"fmt"
	"os"

	"github.com/secnex/dns-server/config"
	"github.com/secnex/dns-server/resolver"
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

	fmt.Printf("Server: %s:%d\n", config.Resolver.Host, config.Resolver.Port)

	resolver := resolver.NewDNSResolver(config)
	resolver.Listen()
}
