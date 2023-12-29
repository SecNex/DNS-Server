package main

import (
	"github.com/secnex/dns-server/resolver"
)

func main() {
	registry := resolver.NewRegistry("0.0.0.0", 53)
	registry.Forwarding.SetForwarding(true, "1.1.1.1")
	domain := resolver.NewDomain("home.lab")
	domain.AddA("@", "192.168.111.250", 300)
	domain.AddA("pi.a", "192.168.111.251", 300)
	domain.AddA("pi.b", "192.168.110.251", 300)
	registry.AddDomain(domain)
	registry.Blocker.AddList("easylist", "https://easylist.to/easylist/easylist.txt")
	registry.Blocker.AddList("StevenBlack", "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts")
	registry.Blocker.Sync()
	registry.Start()
}
