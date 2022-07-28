package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGINT)

	stop := make(chan string)
	var wg sync.WaitGroup

	go func() {
		for {
			s := <-signalChan
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
				log.Println("Signals received")
				close(stop)
				break
			}
		}
	}()

	wg.Add(2)
	go loop100Times(stop, &wg)
	go loop1000Times(stop, &wg)

	wg.Wait()
	log.Println("Complete!")
}

func loop100Times(stop <-chan string, wg *sync.WaitGroup) {
	i := 0
	defer wg.Done()

	for {
		select {
		case <-stop:
			log.Println("loop100Times - quit")
			return
		default:
			if i > 100 {
				return
			}
			log.Println("loop100Times - ", i)
			time.Sleep(50 * time.Millisecond)
			i++
		}
	}
}

func loop1000Times(stop <-chan string, wg *sync.WaitGroup) {
	i := 0
	defer wg.Done()

	for {
		select {
		case <-stop:
			log.Println("loop1000Times - quit")
			return
		default:
			if i > 1000 {
				return
			}
			log.Println("loop1000Times - ", i)
			time.Sleep(50 * time.Millisecond)
			i++
		}
	}
}
