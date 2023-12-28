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

func (r *DnsRegistry) GetDomain(name string) (DnsDomain, bool) {
	for _, d := range r.Domains {
		if d.Name == name {
			return d, true
		}
	}
	return DnsDomain{}, false
}

// UpdateDomain updates the domain with the given name.
func (r *DnsRegistry) UpdateDomain(name string, d DnsDomain) {
	for i, domain := range r.Domains {
		if domain.Name == name {
			r.Domains[i] = d
		}
	}
}
