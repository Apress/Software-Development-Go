package dns

import (
	"errors"
	"golang.org/x/net/dns/dnsmessage"
	"log"
	"net"
)

var (
	errNotSupported = errors.New("type not supported")
)

type Server struct {
	resolver *DNSResolver
	conn     *net.UDPConn
}

func NewServer(conn *net.UDPConn, resolver *DNSResolver) *Server {
	return &Server{
		resolver: resolver,
		conn:     conn,
	}
}

func (s *Server) readRequest() (dnsmessage.Message, *net.UDPAddr, error) {
	buf := make([]byte, 1024)
	_, addr, err := s.conn.ReadFromUDP(buf)
	if err != nil {
		return dnsmessage.Message{}, nil, err
	}
	var msg dnsmessage.Message
	err = msg.Unpack(buf)
	if err != nil {
		return dnsmessage.Message{}, nil, err
	}
	return msg, addr, nil
}

func (s *Server) Start() {
	for {
		err := s.handleRequest()
		if err != nil {
			log.Println(err)
		}
	}
}

func (s *Server) handleRequest() error {
	msg, clientAddr, err := s.readRequest()
	if err != nil {
		return err
	}

	log.Println("Questions : ", msg.Questions)
	rMsg, err := s.resolver.ResolveDNS(msg)
	if err != nil {
		s.sendResponseWithError(clientAddr, msg, err)
		return err
	}
	log.Println("Answers : ", rMsg.Answers)
	return s.sendResponse(clientAddr, rMsg)
}

func (s *Server) sendResponseWithError(clientAddr *net.UDPAddr, msg dnsmessage.Message, err error) {
	switch err {
	case errNotSupported:
		msg.Header.RCode = dnsmessage.RCodeNotImplemented
	default:
		msg.Header.RCode = dnsmessage.RCodeRefused
	}
	err = s.sendResponse(clientAddr, msg)
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) sendResponse(addr *net.UDPAddr, message dnsmessage.Message) error {
	packed, err := message.Pack()
	if err != nil {
		return err
	}
	_, err = s.conn.WriteToUDP(packed, addr)
	return err
}
