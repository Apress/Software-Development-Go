package main

import (
	"k8s.io/klog/v2"
	"log"

	"github.com/google/cadvisor/client"
	info "github.com/google/cadvisor/info/v1"
)

func streamingClientExample(url string) {
	streamingClient, err := client.NewClient("http://localhost:8080/")
	if err != nil {
		klog.Errorf("tried to make client and got error %v", err)
		return
	}
	einfo := make(chan *info.Event)
	go func() {
		err = streamingClient.EventStreamingInfo(url, einfo)
		if err != nil {
			log.Fatalln("got error retrieving event info: ", err)
			return
		}
	}()
	for ev := range einfo {
		log.Println("streaming einfo: ", ev)
	}
}

func main() {
	streamingClientExample("?stream=true&all_events=true&subcontainers=true")
}
