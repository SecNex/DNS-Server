package resolver

import (
	"log"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/secnex/dns-server/config"
	"golang.org/x/net/dns/dnsmessage"
)

type DNSResolver struct {
	Host    string
	Port    int
	Backend *Backend
}

type DNSDomain struct {
	Id          uuid.UUID
	Name        string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type DNSRecord struct {
	Id        uuid.UUID
	Domain    uuid.UUID
	Name      string
	Type      string
	TTL       int
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewDNSResolver(c config.ConfigFile) DNSResolver {
	backend, err := NewBackend(c.Backend)
	if err != nil {
		log.Fatal(err)
	}

	return DNSResolver{
		Host:    c.Resolver.Host,
		Port:    c.Resolver.Port,
		Backend: backend,
	}
}

func (r DNSResolver) Listen() {
	cnx, err := net.ListenUDP("udp", &net.UDPAddr{Port: r.Port})
	if err != nil {
		log.Fatal(err)
	}
	defer cnx.Close()

	for {
		r.handleConnection(cnx)
	}
}

func (r DNSResolver) handleConnection(cnx *net.UDPConn) {
	buffer := make([]byte, 1024)
	n, addr, err := cnx.ReadFromUDP(buffer)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Received %d bytes from %s\n", n, addr)
	var msg dnsmessage.Message
	err = msg.Unpack(buffer[:n])
	if err != nil {
		log.Fatal(err)
	}

	answers := r.Resolve(msg.Questions)
	if len(answers) == len(msg.Questions) {
		msg.Answers = answers
	} else {
		log.Printf("Not all questions resolved\n")
	}
	msg.Response = true
	msg.Authoritative = true
	msg.RecursionAvailable = true
	msg.RecursionDesired = true
	msg.RCode = dnsmessage.RCodeSuccess

	response, err := msg.Pack()
	if err != nil {
		log.Fatal(err)
	}

	_, err = cnx.WriteToUDP(response, addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Sent %d bytes to %s\n", len(response), addr)
}

func (r DNSResolver) Resolve(q []dnsmessage.Question) []dnsmessage.Resource {
	var resources []dnsmessage.Resource
	for _, question := range q {
		domain, record := GetQuery(question)
		d := r.GetDomain(domain)
		r, ok := r.GetRecord(d.Id, record)
		if ok {
			ip := net.ParseIP(r.Value)
			ipBytes := ip.To4()
			resources = append(resources, dnsmessage.Resource{
				Header: dnsmessage.ResourceHeader{
					Name:  question.Name,
					Type:  dnsmessage.TypeA,
					Class: dnsmessage.ClassINET,
					TTL:   uint32(r.TTL),
				},
				Body: &dnsmessage.AResource{
					A: [4]byte{ipBytes[0], ipBytes[1], ipBytes[2], ipBytes[3]},
				},
			})
		}
	}
	return resources
}

func GetQuery(q dnsmessage.Question) (string, string) {
	sub := IsSubdomain(q.Name.String())
	log.Printf("Subdomain: %t\n", sub)
	if sub {
		return GetDomainName(q.Name.String()), GetRecordName(q.Name.String())
	} else {
		return q.Name.String(), "@"
	}
}

func (r DNSResolver) GetDomain(domain string) DNSDomain {
	query := `SELECT * FROM dns.domains WHERE name = $1`
	domainRows, err := r.Backend.Connection.Query(query, domain)
	if err != nil {
		log.Fatal(err)
	}

	var domains []DNSDomain
	for domainRows.Next() {
		var domain DNSDomain
		err := domainRows.Scan(&domain.Id, &domain.Name, &domain.Description, &domain.Status, &domain.CreatedAt, &domain.UpdatedAt, &domain.DeletedAt)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Domain found: %s\n", domain.Name)
		domains = append(domains, domain)
	}

	if len(domains) == 0 {
		log.Printf("Domain not found: %s\n", domain)
		return DNSDomain{}
	} else {
		return domains[0]
	}
}

func (r DNSResolver) GetRecord(d uuid.UUID, record string) (DNSRecord, bool) {
	log.Printf("Searching for record: %s\n", record)
	// select * from dns.records WHERE "domainId" = 'affe54eb-7f43-4cd8-9242-7c3daab8431b' AND name = 'test'
	query := `SELECT * FROM dns.records WHERE "domainId" = $1 AND name = $2 LIMIT 1`
	recordRows, err := r.Backend.Connection.Query(query, d, record)
	if err != nil {
		log.Fatal(err)
	}

	var records []DNSRecord
	for recordRows.Next() {
		var record DNSRecord
		err := recordRows.Scan(&record.Id, &record.Domain, &record.Name, &record.Type, &record.TTL, &record.Value, &record.CreatedAt, &record.UpdatedAt, &record.DeletedAt)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Record found: %s -> %s\n", record.Name, record.Value)
		records = append(records, record)
	}

	if len(records) == 0 {
		log.Printf("Record not found: %s\n", record)
		return DNSRecord{}, false
	} else {
		return records[0], true
	}
}
