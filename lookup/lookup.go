package lookup

import (
	"encoding/json"
	"os"

	"github.com/rianhasiando/go-dns-server/payload"
)

func LookupRecord(payload *payload.Payload) error {
	var (
		database DNSDatabase
		err      error
	)

	jsonBytes, err := os.ReadFile("database.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonBytes, &database)
	if err != nil {
		return err
	}

	return nil
}
