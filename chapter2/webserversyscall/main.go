package main

import (
	"log"
	"net"
	"syscall"
)

const (
	host    = "127.0.0.1"
	port    = 8888
	message = "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html; charset=utf-8\r\n" +
		"Content-Length: 25\r\n" +
		"\r\n" +
		"Server with syscall"
)

func startServer(host string, port int) (int, error) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatal("error (listen) : ", err)
	}

	srv := &syscall.SockaddrInet4{Port: port}
	addrs, err := net.LookupHost(host)
	if err != nil {
		log.Fatal("error (lookup) : ", err)
	}
	for _, addr := range addrs {
		ip := net.ParseIP(addr).To4()
		copy(srv.Addr[:], ip)
		if err = syscall.Bind(fd, srv); err != nil {
			log.Fatal("error (bind) : ", err)
		}
	}
	if err = syscall.Listen(fd, syscall.SOMAXCONN); err != nil {
		log.Fatal("error (listening) : ", err)
	} else {
		log.Println("Listening on ", host, ":", port)
	}
	if err != nil {
		log.Fatal("error (port listening) : ", err)
	}
	return fd, nil
}

func main() {
	fd, err := startServer(host, port)
	if err != nil {
		log.Fatal("error (startServer) : ", err)
	}

	for {
		cSock, cAddr, err := syscall.Accept(fd)

		if err != nil {
			log.Fatal("error (accept) : ", err)
		}

		go func(clientSocket int, clientAddress syscall.Sockaddr) {
			err := syscall.Sendmsg(clientSocket, []byte(message), []byte{}, clientAddress, 0)

			if err != nil {
				log.Fatal("error (send) : ", err)
			}

			syscall.Close(clientSocket)
		}(cSock, cAddr)
	}
}
