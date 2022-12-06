package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const (
	ICMPv4     = 1 // ICMP for IPv4
	ListenAddr = "0.0.0.0"
)

func main() {
	addr := "golang.org"
	dst, dur, err := Ping(addr)
	if err != nil {
		panic(err)
	}
	log.Printf("Ping %s (%s): %s\n", addr, dst, dur)
}

//Ping provide Linux like ping feature
func Ping(addr string) (*net.IPAddr, time.Duration, error) {
	// Listen for ICMP reply
	c, err := icmp.ListenPacket("ip4:icmp", ListenAddr)
	if err != nil {
		return nil, 0, err
	}
	defer c.Close()

	// Resolve DNS to get real IP of the target
	dst, err := net.ResolveIPAddr("ip4", addr)
	if err != nil {
		panic(err)
		return nil, 0, err
	}

	// Prepare new ICMP message
	m := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: []byte(""),
		},
	}

	// Marshal the data
	b, err := m.Marshal(nil)
	if err != nil {
		return dst, 0, err
	}

	// Start measuring time
	start := time.Now()

	// Send ICMP packet now
	n, err := c.WriteTo(b, dst)

	// Check for error
	if err != nil {
		return dst, 0, err
	} else if n != len(b) {
		return dst, 0, fmt.Errorf("got %v; want %v", n, len(b))
	}

	// Allocate 1500 byte for reading response
	reply := make([]byte, 1500)

	// Set deadline of 1 minute
	err = c.SetReadDeadline(time.Now().Add(1 * time.Minute))
	if err != nil {
		return dst, 0, err
	}

	// Read from the connection
	n, peer, err := c.ReadFrom(reply)
	if err != nil {
		return dst, 0, err
	}
	duration := time.Since(start)

	// Use ParseMessage to parsed the bytes received
	rm, err := icmp.ParseMessage(ICMPv4, reply[:n])
	if err != nil {
		return dst, 0, err
	}

	// Check for the type of ICMP result
	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		return dst, duration, nil
	default:
		return dst, 0, fmt.Errorf("got %+v from %v; want echo reply", rm, peer)
	}
}
