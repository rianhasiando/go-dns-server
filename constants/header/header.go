package header

type Header struct {
	TransactionID [2]byte

	QueryType QueryType
	Opcode    Opcode
}
