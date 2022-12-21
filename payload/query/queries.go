package query

type Query struct {
	// the index (0-based) of the first byte of this query
	// in the request
	// (usually the first query has FirstByteIdx of 12 (the 13th bytes))
	FirstByteIdx int

	TLD    string // Top Level Domain
	Domain string // Domain name (abc in abc.com)

	QName  string // full domain address, eg. `Domain.TLD`
	QType  QType
	QClass QClass
}

type QType int

// References: https://en.wikipedia.org/wiki/List_of_DNS_record_types
const (
	QTypeA     QType = 1
	QTypeNS    QType = 2
	QTypeCNAME QType = 5
	QTypeMX    QType = 15
	QTypeTXT   QType = 16
	QTypeAAAA  QType = 28
)

type QClass int

const (
	QClassIN  QClass = 1
	QClassAny QClass = 255
)
