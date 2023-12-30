package resolver

import "golang.org/x/net/dns/dnsmessage"

func (d *DnsDomain) AddA(name string, value string, ttl uint32) {
	record := DnsRecord{
		Name:  name,
		Value: value,
		TTL:   ttl,
		// Pref:  pref,
		Type: dnsmessage.TypeA,
	}
	d.A = append(d.A, record)
}

func (d *DnsDomain) AddAAAA(name string, value string, ttl uint32) {
	record := DnsRecord{
		Name:  name,
		Value: value,
		TTL:   ttl,
		// Pref:  pref,
		Type: dnsmessage.TypeAAAA,
	}
	d.AAAA = append(d.AAAA, record)
}

func (d *DnsDomain) AddCNAME(name string, value string, ttl uint32) {
	record := DnsRecord{
		Name:  name,
		Value: value,
		TTL:   ttl,
		// Pref:  pref,
		Type: dnsmessage.TypeCNAME,
	}
	d.CNAME = append(d.CNAME, record)
}

func (d *DnsDomain) AddMX(name string, value string, ttl uint32, pref uint16) {
	record := DnsRecord{
		Name:  name,
		Value: value,
		TTL:   ttl,
		Pref:  pref,
		Type:  dnsmessage.TypeMX,
	}
	d.MX = append(d.MX, record)
}

func (d *DnsDomain) AddNS(name string, value string, ttl uint32) {
	record := DnsRecord{
		Name:  name,
		Value: value,
		TTL:   ttl,
		// Pref:  pref,
		Type: dnsmessage.TypeNS,
	}
	d.NS = append(d.NS, record)
}

func (d *DnsDomain) AddPTR(name string, value string, ttl uint32) {
	record := DnsRecord{
		Name:  name,
		Value: value,
		TTL:   ttl,
		// Pref:  pref,
		Type: dnsmessage.TypePTR,
	}
	d.PTR = append(d.PTR, record)
}

func (d *DnsDomain) AddSOA(name string, value string, ttl uint32) {
	record := DnsRecord{
		Name:  name,
		Value: value,
		TTL:   ttl,
		// Pref:  pref,
		Type: dnsmessage.TypeSOA,
	}
	d.SOA = append(d.SOA, record)
}

func (d *DnsDomain) AddSRV(name string, value string, ttl uint32, pref uint16) {
	record := DnsRecord{
		Name:  name,
		Value: value,
		TTL:   ttl,
		Pref:  pref,
		Type:  dnsmessage.TypeSRV,
	}
	d.SRV = append(d.SRV, record)
}

func (d *DnsDomain) AddTXT(name string, value string, ttl uint32) {
	record := DnsRecord{
		Name:  name,
		Value: value,
		TTL:   ttl,
		// Pref:  pref,
		Type: dnsmessage.TypeTXT,
	}
	d.TXT = append(d.TXT, record)
}
