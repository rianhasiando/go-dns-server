package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"time"
)

func main() {
	connection, err := net.ListenPacket("udp", ":53")
	if err != nil {
		panic(err.Error())
	}

	defer connection.Close()

	for {
		fmt.Println("socket listener started", time.Now().Format("2006-01-02T15:04:05"))
		buffer := make([]byte, 512)
		n, address, err := connection.ReadFrom(buffer)
		if err != nil {
			fmt.Println(err)
			break
		}
		connection.WriteTo([]byte("ASDASDASD"), address)
		fmt.Println(n, address, hex.EncodeToString(buffer))
	}
}
