//example of showing how to use google gopacket project
//code in this example is to capture traffic from local
//network
package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"log"
	"os"
	"strings"
)

const (
	iface = "enp0s31f6"
	sLen  = 65535
	fName = "test.pcap"
)

var (
	handle      *pcap.Handle
	packetCount int = 0
)

func main() {
	// Open output pcap file and write header
	f, _ := os.Create(fName)
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(sLen, layers.LinkTypeEthernet)
	defer f.Close()

	// Open network interface
	handle, err := pcap.OpenLive(iface, sLen, true, -1)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	pSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range pSource.Packets() {
		printPacketInfo(packet)
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount++

		// Limit the number of capture
		if packetCount > 1000 {
			break
		}
	}
}

func printPacketInfo(packet gopacket.Packet) {
	// extract IPv4 layer
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		log.Println(fmt.Sprintf(
			"(%s) Source address : %s, Destination address : %s",
			ip.Protocol,
			ip.SrcIP,
			ip.DstIP))
	}

	// extract directly DNS layer
	dnsLayer := packet.Layer(layers.LayerTypeDNS)
	if dnsLayer != nil {
		dns, _ := dnsLayer.(*layers.DNS)
		log.Println(fmt.Sprintf("(DNS) Payload %s", string(dns.LayerContents())))
	}

	// extract directly UDP layer
	udpLayer := packet.Layer(layers.LayerTypeUDP)
	if udpLayer != nil {
		tcp, _ := udpLayer.(*layers.UDP)
		log.Println(fmt.Sprintf("(UDP) From port %d to %d", tcp.SrcPort, tcp.DstPort))
	}

	// extract directly TCP layer
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		log.Println(fmt.Sprintf("(TCP) From port %d to %d", tcp.SrcPort, tcp.DstPort))
	}

	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		// Search for a string inside the payload
		if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
			//log.Println("HTTP found!")
			log.Println("HTTP Application layer")
			log.Println("----------------------")
			log.Println(fmt.Sprintf("%s", string(applicationLayer.Payload())))
			log.Println("----------------------")
		}
	}

	// Check for errors
	if err := packet.ErrorLayer(); err != nil {
		log.Println("Error decoding some part of the packet:", err)
	}
}
