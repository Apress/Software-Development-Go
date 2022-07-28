package main

import (
	"log"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 3000,
		IP:   net.ParseIP("0.0.0.0"),
	})
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	log.Printf("Listening %s\n", conn.LocalAddr().String())

	for {
		message := make([]byte, 512)
		_, remoteAddr, err := conn.ReadFromUDP(message[:])
		if err != nil {
			panic(err)
		}

		data := strings.TrimSpace(string(message))
		log.Printf("Received: %s from %s\n", data, remoteAddr)
		conn.WriteToUDP([]byte("\n"), remoteAddr)
	}
}
