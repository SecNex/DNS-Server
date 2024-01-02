package resolver

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func NewRandom(len int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len)))
	return int(n.Int64())
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
	parts := strings.Split(q, ".")[:len(strings.Split(q, "."))-1]
	if len(parts) == 2 {
		return "@"
	} else {
		return strings.Join(parts[:len(parts)-2], ".")
	}
}

func GetDomainName(q string) string {
	// record.home.lab. -> home.lab.
	parts := strings.Split(q, ".")[:len(strings.Split(q, "."))-1]
	if len(parts) == 2 {
		return q
	} else {
		return strings.Join(parts[len(parts)-2:], ".") + "."
	}
}

func IsSubdomain(v string) bool {
	parts := strings.Split(v, ".")[:len(strings.Split(v, "."))-1]
	if len(parts) > 2 {
		return true
	} else {
		return false
	}
}
