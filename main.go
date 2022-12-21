package main

import (
	"encoding/hex"
	"log"
	"net"

	"github.com/rianhasiando/go-dns-server/lookup"
	"github.com/rianhasiando/go-dns-server/payload"
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

		request, err := payload.ParseRawRequest(truncatedRawRequest)
		if err != nil {
			log.Println(err)
			break
		}

		log.Printf("request: %+v\n", request)

		err = lookup.LookupRecord(&request)
		if err != nil {
			log.Println(err)
			break
		}

		connection.WriteTo(truncatedRawRequest, clientAddress)
	}
}
