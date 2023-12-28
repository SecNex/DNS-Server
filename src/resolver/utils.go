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

func CheckIfRoot(q string) bool {
	parts := strings.Split(q, ".")
	// Remove the last empty string
	parts = parts[:len(parts)-1]
	return len(parts) == 2
}
