package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	t := net.JoinHostPort("localhost", "3000")
	l, err := net.Listen("tcp", t)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	log.Println("Listening on port 3000")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	len, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading:", err.Error())
	}

	log.Println("Received : ", string(buf))

	conn.Write([]byte(fmt.Sprintf("Message received of length : %d", len)))
	conn.Close()
}
