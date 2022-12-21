package rr

import (
	q "github.com/rianhasiando/go-dns-server/payload/query"
)

type ResourceRecord struct {
	PointerIdx int // pointer to name in payload (0 if not exists)
	Type       q.QType
	Class      q.QClass
	TTL        int    // in seconds
	RDataReal  string // real data (not ready for response yet)
}
