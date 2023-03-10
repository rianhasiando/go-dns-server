package lookup

import (
	"encoding/json"
	"os"

	"github.com/rianhasiando/go-dns-server/payload"
	e "github.com/rianhasiando/go-dns-server/payload/error"
	q "github.com/rianhasiando/go-dns-server/payload/query"
	"github.com/rianhasiando/go-dns-server/payload/rr"
)

func LookupRecord(p *payload.Payload) error {
	var (
		database DNSDatabase
	)

	jsonBytes, err := os.ReadFile("database.json")
	if err != nil {
		return e.ErrDB
	}

	err = json.Unmarshal(jsonBytes, &database)
	if err != nil {
		return e.ErrDB
	}

	classData := DNSClassData{}
	if p.Query.QClass == q.QClassIN {
		classData = database.ListZones[p.Query.Domain]["IN"]
	}

	if p.Query.QType == q.QTypeA {
		for _, result := range classData.A {
			p.ResourceRecords = append(p.ResourceRecords, rr.ResourceRecord{
				PointerIdx: p.Query.FirstByteIdx,
				Type:       p.Query.QType,
				Class:      p.Query.QClass,
				TTL:        result.TTL,
				RDataReal:  result.Value,
			})
		}
	}

	return nil
}
