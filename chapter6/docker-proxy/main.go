package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"syscall"
)

type handler struct {
	socket string
}

const (
	errCode     = 500
	dsocket     = "/var/run/docker.sock"
	unix        = "unix"
	proxySocket = "/tmp/docker.sock"
)

func writeError(writer http.ResponseWriter, status int, err error) {
	msg := fmt.Sprintf("%v", err)
	log.Print(msg)
	writer.WriteHeader(status)
	writer.Write([]byte(msg))
}

func (h *handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	conn, err := net.DialUnix(unix, nil, &net.UnixAddr{h.socket, unix})
	if err != nil {
		writeError(response, errCode, err)
		return
	}
	defer conn.Close()

	// dump the request into byte[] including the body
	b, err := httputil.DumpRequest(request, true)
	if err != nil {
		log.Fatalf("Error : %v", err)
		return
	}

	log.Printf("[Request] : %s", string(b))

	// write/forward the request to the real Docker socket
	err = request.Write(conn)
	if err != nil {
		writeError(response, errCode, err)
		return
	}

	// get the response from the real Docker socket
	resp, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		writeError(response, errCode, err)
		return
	}
	defer resp.Body.Close()
	for k, v := range resp.Header {
		response.Header()[k] = v
	}
	response.WriteHeader(resp.StatusCode)

	reader := bufio.NewReader(resp.Body)
	for {
		line, _, err := reader.ReadLine()

		//must check for io.EOF as this does not mean
		//problem it only means that Docker has stop
		//sending data. Return only.
		if err == io.EOF {
			return
		} else if err != nil {
			log.Fatalf("Error : %v", err)
			return
		}

		// write the response back to the caller
		response.Write(line)
		log.Printf("[Response] : %s", line)
	}
}

func main() {
	in := flag.String("in", proxySocket, "Proxy docker socket")
	flag.Parse()

	sock, err := net.Listen("unix", *in)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	dhandler := &handler{dsocket}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGABRT, syscall.SIGTERM)
	go func(c chan os.Signal) {
		sig := <-sigChan
		log.Printf("Got signal %v: exiting", sig)
		sock.Close()
		os.Exit(1)
	}(sigChan)

	log.Printf("Listening on %s for Docker commands", proxySocket)
	err = http.Serve(sock, dhandler)
	if err != nil {
		log.Fatalf("Failed to serve http: %v", err)
		return
	}
}
