package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/netpoll"
)

func main() {
	netpoll.SetLoadBalance(netpoll.Random)
	listener, err := netpoll.CreateListener("tcp", "127.0.0.1:8000")
	if err != nil {
		panic("Failure to create listener")
	}

	var opts = []netpoll.Option{
		netpoll.WithIdleTimeout(1 * time.Second),
		netpoll.WithIdleTimeout(10 * time.Minute),
	}
	eventLoop, err := netpoll.NewEventLoop(echoHandler, opts...)
	if err != nil {
		panic("Failure to create netpoll")
	}
	err = eventLoop.Serve(listener)
	if err != nil {
		panic("Failure to run netpoll")
	}
}

//echoHandler handles incoming request
func echoHandler(ctx context.Context, connection netpoll.Connection) error {
	reader := connection.Reader()
	bts, err := reader.Next(reader.Len())
	if err != nil {
		log.Println("error reading data")
		return err
	}
	log.Println(fmt.Sprintf("Data: %s", string(bts)))

	connection.Write([]byte("-> " + string(bts)))
	return connection.Writer().Flush()
}
