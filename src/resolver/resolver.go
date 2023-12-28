package resolver

type DnsForwarding struct {
	Enabled bool   `json:"enabled"`
	Server  string `json:"server"`
}

type DnsRegistry struct {
	Host        string          `json:"host"`
	Port        int             `json:"port"`
	Domains     []DnsDomain     `json:"domains"`
	Forwarding  DnsForwarding   `json:"forwarding"`
	RootServers []DnsRootServer `json:"rootServers"`
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
	Name  string `json:"name"`
	Value string `json:"value"`
	TTL   int    `json:"ttl"`
	Pref  uint16 `json:"pref"`
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
	}
}

func newForwarding(enabled bool, server string) DnsForwarding {
	return DnsForwarding{
		Enabled: enabled,
		Server:  server,
	}
}

func (f *DnsForwarding) SetForwarding(enabled bool, server string) {
	f.Enabled = enabled
	f.Server = server
}
