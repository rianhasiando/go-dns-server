package payload

import (
	"errors"
	"fmt"

	"github.com/rianhasiando/go-dns-server/payload/header"
	"github.com/rianhasiando/go-dns-server/payload/query"
)

type Payload struct {
	Header header.Header
	Query  query.Query
}

func ParseRawRequest(rawRequest []byte) (Payload, error) {
	request := Payload{}

	if len(rawRequest) < 12 {
		return request, errors.New("header length must be at least 12 bytes")
	}

	// transaction ID  (2 bytes)
	request.Header.TransactionID = [2]byte{rawRequest[0], rawRequest[1]}

	// 2 bytes of flags
	flags := [2]byte{rawRequest[2], rawRequest[3]}
	request.Header.QueryType = header.QueryType(
		(0b10000000 & flags[0]) >> 7,
	)
	request.Header.Opcode = header.Opcode(
		(0b01111000 & flags[0]) >> 3,
	)
	request.Header.Truncation = ((0b00000010 & flags[0]) >> 1) == 1
	request.Header.RecursionDesired = (0b00000001 & flags[0]) == 1
	request.Header.RecursionAvailable = ((0b10000000 & flags[1]) >> 7) == 1
	request.Header.Reserved = int((0b01000000 & flags[1]) >> 6)
	request.Header.AuthenticData = ((0b00100000 & flags[1]) >> 5) == 1
	request.Header.CheckingDisabled = ((0b00010000 & flags[1]) >> 4) == 1
	request.Header.ResponseCode = header.ResponseCode(
		(0b00001111 & flags[1]),
	)

	// 2 bytes of QDCOUNT
	request.Header.QuestionCount = int(rawRequest[4])<<8 + int(rawRequest[5])

	// 2 bytes of ANCOUNT
	request.Header.AnswerCount = int(rawRequest[6])<<8 + int(rawRequest[7])

	// 2 bytes of NSCOUNT
	request.Header.NameServerCount = int(rawRequest[8])<<8 + int(rawRequest[9])

	// 2 bytes of ARCOUNT
	request.Header.AdditionalRecordsCount = int(rawRequest[10])<<8 + int(rawRequest[11])

	currentPointerIdx := 12
	// parsing the queries
	// typically one request only have 1 question
	// so we need to only process one query (if given)
	if request.Header.QuestionCount == 1 && len(rawRequest) >= currentPointerIdx+1 {

		numCharDomain := uint(rawRequest[currentPointerIdx])

		if len(rawRequest) >= currentPointerIdx+int(numCharDomain)+1 {
			currentPointerIdx += 1
			request.Query.Domain = string(rawRequest[currentPointerIdx : currentPointerIdx+int(numCharDomain)])
			currentPointerIdx += int(numCharDomain - 1)
		}

		numCharTLD := uint(0)
		currentPointerIdx += 1
		if len(rawRequest) >= currentPointerIdx+1 {
			numCharTLD = uint(rawRequest[currentPointerIdx])

			currentPointerIdx += 1
			if len(rawRequest) >= currentPointerIdx+int(numCharTLD)+1 {
				request.Query.TLD = string(rawRequest[currentPointerIdx : currentPointerIdx+int(numCharTLD)])
				currentPointerIdx += int(numCharTLD - 1)
			}
		}

		request.Query.QName = fmt.Sprintf("%s.%s",
			request.Query.Domain,
			request.Query.TLD,
		)

		// after the domain specs, the next byte is 0x00 (null character)
		// we just skip this null character because it's not needed anyway
		currentPointerIdx += 2

		// make sure we have 2 bytes more for qtype
		if len(rawRequest) >= currentPointerIdx+2 {
			t := int(rawRequest[currentPointerIdx]) << 8
			currentPointerIdx += 1
			t += int(rawRequest[currentPointerIdx])
			request.Query.QType = query.QType(t)
		}

		currentPointerIdx += 1
		// make sure we have 2 bytes more for qclass
		if len(rawRequest) >= currentPointerIdx+2 {
			c := int(rawRequest[currentPointerIdx]) << 8
			currentPointerIdx += 1
			c += int(rawRequest[currentPointerIdx])
			request.Query.QClass = query.QClass(c)
		}
	}

	// here we don't parse additional request for simplicity

	return request, nil
}
