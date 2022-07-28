package main

import (
	"log"
	"math/rand"
	"syscall"
	"time"
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func randomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	log.Printf("Starting app")

	dirPath := "/tmp/" + randomString(8)
	if err := syscall.Mkdir(dirPath, 0600); err != nil {
		log.Printf("Failed creating directory: %v", err)
		return
	}
	log.Printf("Directory %s created successfully", dirPath)

	// Trying to run non whitelisted syscall
	log.Println("Trying to get current working directory")
	wd, err := syscall.Getwd()
	if err != nil {
		log.Printf("Failed getting current working directory: %v", err)
		return
	}
	log.Printf("Current working directory is: %s", wd)

}
