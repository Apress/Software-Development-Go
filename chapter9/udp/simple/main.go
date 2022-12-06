package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println(qotd("djxmmx.net:17"))
}

//qotd read from quote-of-day server
func qotd(s string) string {
	udpAddr, err := net.ResolveUDPAddr("udp", s)
	if err != nil {
		println("Error Resolving UDP Address:", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)

	buffer := make([]byte, 1024)
	conn.Write([]byte("\n"))
	conn.Read(buffer[0:])
	conn.Close()

	return string(buffer)
}
