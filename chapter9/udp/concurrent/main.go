package main

import (
	"fmt"
	"log"
	"net"
	"runtime"
)

func listen(i int, connection *net.UDPConn, quit chan struct{}) {
	buffer := make([]byte, 512)
	for {
		_, remote, err := connection.ReadFromUDP(buffer)
		if err != nil {
			break
		}

		log.Println(fmt.Sprintf("(Routine %d) Connection from %s - %s", i, remote.String(),
			string(buffer)))
		connection.WriteToUDP([]byte("\n"), remote)
	}
	quit <- struct{}{}
}

func main() {
	addr := net.UDPAddr{
		Port: 3333,
	}
	connection, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	quit := make(chan struct{})
	for i := 0; i < runtime.NumCPU(); i++ {
		log.Println("Spinning up UDP server - ", i)
		id := i
		go listen(id, connection, quit)
	}
	<-quit
}
