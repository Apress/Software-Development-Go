package main

import (
	"log"
	"runtime"
)

func main() {
	log.Println("Hello, from inside Docker image")
	log.Println("Build using Go version ", runtime.Version())
}
