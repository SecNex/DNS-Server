package resolver

import (
	"fmt"
	"net"
	"strings"

	"golang.org/x/net/dns/dnsmessage"
)

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
	// Extract the domain name from the question. -> record.home.lab -> home.lab
	parts := strings.Split(name, ".")
	domainName := strings.Join(parts[len(parts)-2:], ".")
	for _, d := range r.Domains {
		if d.Name == domainName {
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

// GetRecord returns the record with the given name and type.
func (d *DnsDomain) GetRecord(name string, t dnsmessage.Type) (DnsRecord, bool) {
	switch t {
	case dnsmessage.TypeA:
		for _, record := range d.A {
			if record.Name == name {
				return record, true
			}
		}
	case dnsmessage.TypeAAAA:
		for _, record := range d.AAAA {
			if record.Name == name {
				return record, true
			}
		}
	case dnsmessage.TypeCNAME:
		for _, record := range d.CNAME {
			if record.Name == name {
				return record, true
			}
		}
	case dnsmessage.TypeMX:
		for _, record := range d.MX {
			if record.Name == name {
				return record, true
			}
		}
	case dnsmessage.TypeNS:
		for _, record := range d.NS {
			if record.Name == name {
				return record, true
			}
		}
	case dnsmessage.TypePTR:
		for _, record := range d.PTR {
			if record.Name == name {
				return record, true
			}
		}
	case dnsmessage.TypeSOA:
		for _, record := range d.SOA {
			if record.Name == name {
				return record, true
			}
		}
	case dnsmessage.TypeSRV:
		for _, record := range d.SRV {
			if record.Name == name {
				return record, true
			}
		}
	case dnsmessage.TypeTXT:
		for _, record := range d.TXT {
			if record.Name == name {
				return record, true
			}
		}
	}
	return DnsRecord{}, false
}

// CreateAnswer creates an answer for the given question.
func (r *DnsRecord) CreateAnswer(q dnsmessage.Question) dnsmessage.Resource {
	fmt.Printf("Creating answer for %s (%s) -> %s\n", q.Name.String(), q.Type, r.Value)
	switch q.Type {
	case dnsmessage.TypeA:
		ip := net.ParseIP(r.Value)
		ip = ip.To4()
		return dnsmessage.Resource{
			Header: dnsmessage.ResourceHeader{
				Name:  q.Name,
				Type:  dnsmessage.TypeA,
				Class: dnsmessage.ClassINET,
				TTL:   r.TTL,
			},
			Body: &dnsmessage.AResource{
				A: [4]byte{ip[0], ip[1], ip[2], ip[3]},
			},
		}
	case dnsmessage.TypeAAAA:
		ip := net.ParseIP(r.Value)
		ip = ip.To16()
		return dnsmessage.Resource{
			Header: dnsmessage.ResourceHeader{
				Name:  q.Name,
				Type:  dnsmessage.TypeAAAA,
				Class: dnsmessage.ClassINET,
				TTL:   r.TTL,
			},
			Body: &dnsmessage.AAAAResource{
				AAAA: [16]byte{
					ip[0], ip[1], ip[2], ip[3],
					ip[4], ip[5], ip[6], ip[7],
					ip[8], ip[9], ip[10], ip[11],
					ip[12], ip[13], ip[14], ip[15],
				},
			},
		}
	case dnsmessage.TypeCNAME:
		return dnsmessage.Resource{
			Header: dnsmessage.ResourceHeader{
				Name:  q.Name,
				Type:  dnsmessage.TypeCNAME,
				Class: dnsmessage.ClassINET,
				TTL:   r.TTL,
			},
			Body: &dnsmessage.CNAMEResource{CNAME: dnsmessage.MustNewName(r.Value)},
		}
	case dnsmessage.TypeMX:
		return dnsmessage.Resource{
			Header: dnsmessage.ResourceHeader{
				Name:  q.Name,
				Type:  dnsmessage.TypeMX,
				Class: dnsmessage.ClassINET,
				TTL:   r.TTL,
			},
			Body: &dnsmessage.MXResource{
				MX:   dnsmessage.MustNewName(r.Value),
				Pref: r.Pref,
			},
		}
	case dnsmessage.TypeNS:
		return dnsmessage.Resource{
			Header: dnsmessage.ResourceHeader{
				Name:  q.Name,
				Type:  dnsmessage.TypeNS,
				Class: dnsmessage.ClassINET,
				TTL:   r.TTL,
			},
			Body: &dnsmessage.NSResource{NS: dnsmessage.MustNewName(r.Value)},
		}
	case dnsmessage.TypePTR:
		return dnsmessage.Resource{
			Header: dnsmessage.ResourceHeader{
				Name:  q.Name,
				Type:  dnsmessage.TypePTR,
				Class: dnsmessage.ClassINET,
				TTL:   r.TTL,
			},
			Body: &dnsmessage.PTRResource{PTR: dnsmessage.MustNewName(r.Value)},
		}
	case dnsmessage.TypeTXT:
		return dnsmessage.Resource{
			Header: dnsmessage.ResourceHeader{
				Name:  q.Name,
				Type:  dnsmessage.TypeTXT,
				Class: dnsmessage.ClassINET,
				TTL:   r.TTL,
			},
			Body: &dnsmessage.TXTResource{TXT: []string{r.Value}},
		}
	}
	return dnsmessage.Resource{}
}
