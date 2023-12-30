package resolver

import (
	"encoding/json"
	"fmt"
	"net"

	"golang.org/x/net/dns/dnsmessage"
)

type DnsForwarding struct {
	Enabled bool      `json:"enabled"`
	Server  string    `json:"server"`
	Client  DnsClient `json:"client"`
}

type DnsRegistry struct {
	Host        string          `json:"host"`
	Port        int             `json:"port"`
	Domains     []DnsDomain     `json:"domains"`
	Forwarding  DnsForwarding   `json:"forwarding"`
	RootServers []DnsRootServer `json:"rootServers"`
	Blocker     Blocker         `json:"blocker"`
}

type DnsDomain struct {
	Name  string      `json:"name"`
	A     []DnsRecord `json:"a"`
	AAAA  []DnsRecord `json:"aaaa"`
	CNAME []DnsRecord `json:"cname"`
	MX    []DnsRecord `json:"mx"`
	NS    []DnsRecord `json:"ns"`
	PTR   []DnsRecord `json:"ptr"`
	SOA   []DnsRecord `json:"soa"`
	SRV   []DnsRecord `json:"srv"`
	TXT   []DnsRecord `json:"txt"`
}

type DnsRecord struct {
	Name  string          `json:"name"`
	Value string          `json:"value"`
	TTL   uint32          `json:"ttl"`
	Pref  uint16          `json:"pref"`
	Type  dnsmessage.Type `json:"type"`
}

type DnsRootServer struct {
	Host string `json:"host"`
	IPv4 string `json:"ipv4"`
	IPv6 string `json:"ipv6"`
}

var RootServers = []DnsRootServer{
	{
		Host: "a.root-servers.net",
		IPv4: "198.41.0.4",
		IPv6: "2001:503:ba3e::2:30",
	},
	{
		Host: "b.root-servers.net",
		IPv4: "170.247.170.2",
		IPv6: "2001:500:84::b",
	},
	{
		Host: "c.root-servers.net",
		IPv4: "192.33.4.12",
		IPv6: "2001:500:2::c",
	},
	{
		Host: "d.root-servers.net",
		IPv4: "199.7.91.13",
		IPv6: "2001:500:2d::d",
	},
	{
		Host: "e.root-servers.net",
		IPv4: "192.203.230.10",
		IPv6: "2001:500:a8::e",
	},
	{
		Host: "f.root-servers.net",
		IPv4: "192.5.5.241",
		IPv6: "2001:500:2f::f",
	},
	{
		Host: "g.root-servers.net",
		IPv4: "192.112.36.4",
		IPv6: "2001:500:12::d0d",
	},
	{
		Host: "h.root-servers.net",
		IPv4: "198.97.190.53",
		IPv6: "2001:500:1::53",
	},
	{
		Host: "i.root-servers.net",
		IPv4: "192.36.148.17",
		IPv6: "2001:7fe::53",
	},
	{
		Host: "j.root-servers.net",
		IPv4: "192.58.128.30",
		IPv6: "2001:503:c27::2:30",
	},
	{
		Host: "k.root-servers.net",
		IPv4: "193.0.14.129",
		IPv6: "2001:7fd::1",
	},
	{
		Host: "l.root-servers.net",
		IPv4: "199.7.83.42",
		IPv6: "2001:500:9f::42",
	},
	{
		Host: "m.root-servers.net",
		IPv4: "202.12.27.33",
		IPv6: "2001:dc3::35",
	},
}

func NewRegistry(host string, port int) DnsRegistry {
	return DnsRegistry{
		Host:        host,
		Port:        port,
		Domains:     []DnsDomain{},
		Forwarding:  newForwarding(false, "8.8.8.8"),
		RootServers: RootServers,
		Blocker:     NewBlocker(),
	}
}

func newForwarding(enabled bool, server string) DnsForwarding {
	return DnsForwarding{
		Enabled: enabled,
		Server:  server,
		Client:  NewClient(server, 53),
	}
}

func (r *DnsRegistry) ConfigJson() string {
	j, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(j)
}

func (r *DnsRegistry) handleQuery(cnx *net.UDPConn, addr *net.UDPAddr, buf []byte) {
	var msg dnsmessage.Message
	err := msg.Unpack(buf)
	if err != nil {
		panic(err)
	}
	var response dnsmessage.Message
	response.Header.ID = msg.Header.ID
	response.Header.Response = true
	response.Header.RecursionDesired = true
	response.Header.RecursionAvailable = true
	response.Header.Authoritative = false
	response.Questions = msg.Questions
	response.Answers = []dnsmessage.Resource{}
	response.Authorities = []dnsmessage.Resource{}
	response.Additionals = []dnsmessage.Resource{}
	for _, question := range msg.Questions {
		// Remove trailing dot
		queryName := question.Name.String()
		if queryName[len(queryName)-1] == '.' {
			queryName = queryName[:len(queryName)-1]
		}
		if !r.Blocker.IsBlocked(queryName) {
			// Check if the domain is configured
			d, ok := r.GetDomain(queryName)
			fmt.Printf("Query for %s (%s)\n", queryName, question.Type)
			if ok {
				recordQueryName := GetRecordName(queryName)
				record, ok := d.GetRecord(recordQueryName, question.Type)
				fmt.Printf("Record: %s\n", record.Name)
				// If record type of question of A or AAAA but not found, check for CNAME
				if (question.Type == dnsmessage.TypeA || question.Type == dnsmessage.TypeAAAA) && (record.Type == dnsmessage.TypeCNAME) {
					question.Type = dnsmessage.TypeCNAME
				}
				if ok {
					response.Answers = append(response.Answers, record.CreateAnswer(question))
				} else {
					response.Answers = []dnsmessage.Resource{}
				}
				ns, ok := d.GetRecord("@", dnsmessage.TypeNS)
				if ok {
					authority := ns.CreateAuthority(question)
					response.Authorities = append(response.Authorities, authority)
					response.Authoritative = true
				}
			} else {
				if r.Forwarding.Enabled {
					fmt.Printf("Forwarding query for %s (%s) -> %s\n", question.Name.String(), question.Type, r.Forwarding.Server)
					answer, err := r.Forwarding.Client.Query(question)
					if err != nil {
						panic(err)
					}
					response.Answers = append(response.Answers, answer.Answers...)
					response.Authorities = append(response.Authorities, answer.Authorities...)
					response.Additionals = append(response.Additionals, answer.Additionals...)
				}
			}
		} else {
			fmt.Printf("Blocked query to %s\n", question.Name.String())
		}
	}
	buf, err = response.Pack()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		panic(err)
	}
	_, err = cnx.WriteToUDP(buf, addr)
	if err != nil {
		panic(err)
	}
}

func (f *DnsForwarding) SetForwarding(enabled bool, server string) {
	f.Enabled = enabled
	f.Server = server
	f.Client = NewClient(server, 53)
}

func (f *DnsForwarding) GetForwarding() DnsForwarding {
	return *f
}
