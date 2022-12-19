package header

type Opcode int

const (
	StandardQuery Opcode = 0
	InverseQuery  Opcode = 1
)
