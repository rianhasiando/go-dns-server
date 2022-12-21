package query

type Query struct {
	TLD    string // Top Level Domain
	Domain string // Domain name (abc in abc.com)

	QName  string // full domain address, eg. `Domain.TLD`
	QType  QType
	QClass QClass
}

type QType int

// References: https://en.wikipedia.org/wiki/List_of_DNS_record_types
const (
	QTypeA     = 1
	QTypeNS    = 2
	QTypeCNAME = 5
	QTypeMX    = 15
	QTypeTXT   = 16
	QTypeAAAA  = 28
)

type QClass int

const (
	QClassIN  = 1
	QClassAny = 255
)
