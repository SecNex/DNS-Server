package resolver

import (
	"fmt"
	"net"
)

func (r *DnsRegistry) Start() {
	cnx, err := newListener(r.Host, r.Port)
	if err != nil {
		panic(err)
	}
	defer cnx.Close()

	fmt.Printf("Listening on %s:%d\n", r.Host, r.Port)

	buf := make([]byte, 512)
	for {
		n, addr, err := cnx.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}

		go func(cnx *net.UDPConn, addr *net.UDPAddr, buf []byte) {
			r.handleQuery(cnx, addr, buf[:n])
		}(cnx, addr, buf)
	}
}

func newListener(host string, port int) (*net.UDPConn, error) {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(host),
	}
	return net.ListenUDP("udp", &addr)
}
