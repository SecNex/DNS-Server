package resolver

import (
	"fmt"
	"net"

	"golang.org/x/net/dns/dnsmessage"
)

type DnsClient struct {
	Host       string
	Port       int
	Connection *net.UDPConn
}

func NewClient(host string, port int) DnsClient {
	serverAddr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(host),
	}
	cnx, err := net.DialUDP("udp", nil, &serverAddr)
	if err != nil {
		panic(err)
	}
	return DnsClient{
		Host:       host,
		Port:       port,
		Connection: cnx,
	}
}

func (c *DnsClient) Query(q dnsmessage.Question) (dnsmessage.Message, error) {
	msg := dnsmessage.Message{
		Header: dnsmessage.Header{
			RecursionDesired: true,
		},
		Questions: []dnsmessage.Question{q},
	}

	buf, err := msg.Pack()
	if err != nil {
		panic(err)
	}

	_, err = c.Connection.Write(buf)
	if err != nil {
		panic(err)
	}

	buf = make([]byte, 1024)
	n, err := c.Connection.Read(buf)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		panic(err)
	}
	var response dnsmessage.Message
	err = response.Unpack(buf[:n])
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		panic(err)
	}
	return response, nil
}
