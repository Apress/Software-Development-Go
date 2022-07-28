package main

import (
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	s := spinner.New(spinner.CharSets[0], 100*time.Millisecond)
	s.Color("red")
	s.Prefix = "Processing request : "
	s.Start()
	done := make(chan struct{}, 1)

	//do some processing for 4 seconds
	go func(d chan struct{}) {
		time.Sleep(4 * time.Second)
		d <- struct{}{}
	}(done)

	<-done
	s.Stop()
}
