package payload

import (
	"github.com/rianhasiando/go-dns-server/definition/payload/header"
	"github.com/rianhasiando/go-dns-server/definition/payload/query"
)

type Payload struct {
	Header header.Header
	Query  query.Query
}
