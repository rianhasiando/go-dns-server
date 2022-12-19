package main

import (
	"encoding/hex"
	"log"
	"net"

	"github.com/rianhasiando/go-dns-server/constants/header"
)

func main() {
	connection, err := net.ListenPacket("udp", ":53")
	if err != nil {
		panic(err.Error())
	}

	defer connection.Close()

	for {
		requestBuffer := make([]byte, 512)
		numBytes, clientAddress, err := connection.ReadFrom(requestBuffer)
		if err != nil {
			log.Println(err)
			break
		}

		truncatedRawRequest := requestBuffer[:numBytes]

		log.Println(numBytes, clientAddress, hex.EncodeToString(truncatedRawRequest))

		request, _ := parseRawRequest(truncatedRawRequest)

		log.Println(request)

		connection.WriteTo(truncatedRawRequest, clientAddress)
	}
}

func parseRawRequest(rawRequest []byte) (header.Header, error) {
	request := header.Header{}
	// transaction ID  (2 bytes)
	request.TransactionID = [2]byte{rawRequest[0], rawRequest[1]}

	flags := [2]byte{rawRequest[2], rawRequest[3]}

	request.QueryType = header.QueryType(
		(0b10000000 & flags[0]) >> 7,
	)
	request.Opcode = header.Opcode(
		(0b01111000 & flags[0]) >> 3,
	)

	return request, nil
}
