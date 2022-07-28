package main

import (
	"log"
	"net"
	"time"
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

		//pretend the code is doing some request processing for 10milliseconds
		time.Sleep(10 * time.Millisecond)
		conn.WriteToUDP([]byte("\n"), remoteAddr)
	}
}
