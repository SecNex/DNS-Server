package resolver

func NewDomain(name string) DnsDomain {
	return DnsDomain{
		Name:  name,
		A:     []DnsRecord{},
		AAAA:  []DnsRecord{},
		CNAME: []DnsRecord{},
		MX:    []DnsRecord{},
		NS:    []DnsRecord{},
		PTR:   []DnsRecord{},
		SOA:   []DnsRecord{},
		SRV:   []DnsRecord{},
		TXT:   []DnsRecord{},
	}
}

func (r *DnsRegistry) AddDomain(d DnsDomain) {
	r.Domains = append(r.Domains, d)
}
