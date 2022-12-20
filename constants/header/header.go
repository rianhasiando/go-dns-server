package header

type Header struct {
	TransactionID [2]byte

	QueryType        QueryType
	Opcode           Opcode
	Truncation       bool
	RecursionDesired bool
	Z                int
	ResponseCode     ResponseCode

	QuestionCount          int
	AnswerCount            int
	NameServerCount        int
	AdditionalRecordsCount int
}
