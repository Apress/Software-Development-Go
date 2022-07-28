package main

import (
	"bufio"
	"fmt"
	"net"
)

const (
	host = "google.com"
	port = "80"
)

func main() {
	t := net.JoinHostPort(host, port)

	conn, err := net.Dial("tcp", t)
	if err != nil {
		panic(err)
	}

	req := "GET / \r\nHost: google.com\r\nConnection: close\r\n\r\n"
	conn.Write([]byte(req))

	connReader := bufio.NewReader(conn)
	scanner := bufio.NewScanner(connReader)

	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error", err)
	}
}
