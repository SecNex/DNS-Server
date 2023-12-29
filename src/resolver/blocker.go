package resolver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Blocker struct {
	Lists []BlockList
}

type BlockList struct {
	Name    string
	Url     string
	Status  bool
	Records []BlockListRecord
}

type BlockListRecord struct {
	Raw string
}

func NewBlocker() Blocker {
	return Blocker{
		Lists: []BlockList{},
	}
}

func (b *Blocker) AddList(name string, url string) {
	b.Lists = append(b.Lists, BlockList{
		Name:    name,
		Url:     url,
		Status:  false,
		Records: []BlockListRecord{},
	})
}

func (b *Blocker) GetList(name string) (BlockList, bool) {
	for _, list := range b.Lists {
		if list.Name == name {
			return list, true
		}
	}
	return BlockList{}, false
}

func (b *Blocker) UpdateList(name string, list BlockList) {
	for i, l := range b.Lists {
		if l.Name == name {
			b.Lists[i] = list
		}
	}
}

func (b *Blocker) Sync() {
	for i, list := range b.Lists {
		if !list.Status {
			b.Lists[i].Records = b.FetchList(list.Url)
			b.Lists[i].Status = true
		}
	}
}

func (b *Blocker) FetchList(url string) []BlockListRecord {
	// Fetch the list from the URL
	fmt.Printf("Fetching list from %s\n", url)

	// HTTP-Anfrage an die URL senden
	resp, err := http.Get(url)
	if err != nil {
		// Fehlerbehandlung
		fmt.Printf("Error fetching list: %s\n", err.Error())
		return nil
	}
	defer resp.Body.Close()

	// Antwortkörper lesen
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Fehlerbehandlung
		fmt.Printf("Error reading response body: %s\n", err.Error())
		return nil
	}

	// Liste der BlockListRecord-Objekte erstellen
	records := []BlockListRecord{}
	for _, line := range strings.Split(string(body), "\n") {
		if len(line) > 0 && line[0] != '#' {
			records = append(records, BlockListRecord{
				Raw: line,
			})
		}
	}

	fmt.Printf("Fetched %d records\n", len(records))

	return records
}

func (b *Blocker) IsBlocked(name string) bool {
	for _, list := range b.Lists {
		for _, record := range list.Records {
			if strings.HasSuffix(name, record.Raw) {
				if strings.Contains(record.Raw, "*") {
					return true
				} else {
					if name == record.Raw {
						return true
					}
				}
			}
		}
	}
	return false
}
