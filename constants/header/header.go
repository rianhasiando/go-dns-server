package header

type Header struct {
	TransactionID [2]byte

	QueryType           QueryType
	Opcode              Opcode
	AuthoritativeAnswer bool
	Truncation          bool
	RecursionDesired    bool
	RecursionAvailable  bool
	Reserved            int  // Z bit
	AuthenticData       bool // AD bit
	CheckingDisabled    bool
	ResponseCode        ResponseCode

	QuestionCount          int
	AnswerCount            int
	NameServerCount        int
	AdditionalRecordsCount int
}
