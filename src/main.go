package main

import (
	"os"
	"strconv"

	"github.com/secnex/dns-server/resolver"
)

func main() {
	host := os.Args[1]
	port, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		panic(err)
	}
	forwarding := ""
	if len(os.Args) > 3 {
		forwarding = os.Args[3]
	}
	registry := resolver.NewRegistry(host, int(port))
	registry.Forwarding.SetForwarding(forwarding != "", forwarding)
	registry.Blocker.AddList("StevenBlack", "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts")
	registry.Blocker.Sync()
	registry.Start()
}
