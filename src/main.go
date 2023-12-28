package main

import (
	"encoding/json"
	"fmt"

	"github.com/secnex/dns-server/resolver"
)

func main() {
	registry := resolver.NewRegistry()
	domain := resolver.NewDomain("bytie.lab")
	domain.AddA("@", "10.100.0.1", 300)
	domain.AddCNAME("www", "bytie.lab", 300)
	registry.AddDomain(domain)
	domain.AddAAAA("@", "2001:db8::1", 300)
	registry.UpdateDomain("bytie.lab", domain)
	myResolver, _ := json.Marshal(registry)
	fmt.Println(string(myResolver))
}
