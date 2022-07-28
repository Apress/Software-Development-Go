package main

import (
	"net"
	"network.dns/server/pkg/dns"
)

type DNSConfig struct {
	dnsForwarder string
	port         int
}

func main() {
	dnsConfig := DNSConfig{
		dnsForwarder: "8.8.8.8:53",
		port:         8090,
	}

	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: dnsConfig.port})
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	dnsFwdConn, err := net.Dial("udp", dnsConfig.dnsForwarder)
	if err != nil {
		panic(err)
	}
	defer dnsFwdConn.Close()

	if err != nil {
		panic(err)
	}

	dnsServer := dns.NewServer(conn, dns.NewUDPResolver(dnsFwdConn))
	dnsServer.Start()
}
