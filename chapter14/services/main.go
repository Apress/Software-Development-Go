package main

import (
	"log"
	"time"
)

func main() {
	serviceBDone := make(chan bool, 1)
	alldone := make(chan bool, 1)

	go serviceB(serviceBDone)
	go serviceA(serviceBDone, alldone)

	<-alldone
}

//1st service
func serviceB(serviceBDone chan bool) {
	log.Println("....Starting serviceB")
	for i := 0; i < 10; i++ {
		time.Sleep(50 * time.Millisecond)
	}

	serviceBDone <- true
	log.Println("....Done with serviceB")
}

//2nd service
func serviceA(serviceBDone chan bool, finish chan bool) {
	<-serviceBDone
	log.Println("..Starting serviceA")
	for i := 0; i < 50; i++ {
		time.Sleep(50 * time.Millisecond)
	}
	log.Println("..Done with serviceA")
	finish <- true
}
