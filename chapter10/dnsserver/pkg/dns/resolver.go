package dns

import (
	"golang.org/x/net/dns/dnsmessage"
	"net"
)

type DNSResolver struct {
	fwdConn net.Conn
}

//NewUDPResolver create new DNSResolver implementation
func NewUDPResolver(conn net.Conn) *DNSResolver {
	return &DNSResolver{
		fwdConn: conn,
	}
}

//ResolveDNS resolve the actual DNS query using fwdConn
func (r *DNSResolver) ResolveDNS(msg dnsmessage.Message) (dnsmessage.Message, error) {
	packedMsg, err := msg.Pack()
	if err != nil {
		return dnsmessage.Message{}, err
	}
	_, err = r.fwdConn.Write(packedMsg)
	if err != nil {
		return dnsmessage.Message{}, err
	}
	resBuf := make([]byte, 1024)
	_, err = r.fwdConn.Read(resBuf)
	if err != nil {
		return dnsmessage.Message{}, err
	}

	var resMsg dnsmessage.Message
	err = resMsg.Unpack(resBuf)
	if err != nil {
		return dnsmessage.Message{}, err
	}
	return resMsg, nil
}
