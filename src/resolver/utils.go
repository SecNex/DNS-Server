package resolver

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func newRandom(len int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len)))
	return int(n.Int64())
}

func RandomRootServer(r *DnsRegistry) DnsRootServer {
	len := len(r.RootServers)
	// Random number generator (0 - len)
	random := newRandom(len)
	return r.RootServers[random]
}

func ReverseIP(ip string) string {
	parts := strings.Split(ip, ".")
	reversedParts := make([]string, len(parts))
	for i := 0; i < len(parts); i++ {
		reversedParts[i] = parts[len(parts)-1-i]
	}
	return strings.Join(reversedParts, ".")
}

func GetRecordName(q string) string {
	// Check if the query is home.lab not record.home.lab
	parts := strings.Split(q, ".")
	if len(parts) == 2 {
		return "@"
	} else {
		return strings.Join(parts[:len(parts)-2], ".")
	}
}
