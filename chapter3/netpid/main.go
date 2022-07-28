package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {

	go func() {
		server := &http.Server{}
		l, err := net.Listen("tcp4", ":8080")
		if err != nil {
			log.Fatal(err)
		}
		err = server.Serve(l)
		//
		//http.HandleFunc("/", HelloServer)
		//http.ListenAndServe(":8080", nil)
	}()

	time.Sleep(1050 * time.Millisecond)
	fmt.Println(os.Getpid())
	m, _ := readFdDir(os.Getpid())
	log.Println(m)

	log.Println(establishedTcpConns())
	log.Println(netTcpConns())

	for {
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
