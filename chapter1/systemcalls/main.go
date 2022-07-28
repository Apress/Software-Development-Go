package main

import (
	u "golang.org/x/sys/unix"
	"log"
)

func main() {
	c := make([]byte, 512)

	log.Println("Getpid : ", u.Getpid())
	log.Println("Getpgrp : ", u.Getpgrp())
	log.Println("Getppid : ", u.Getppid())
	log.Println("Gettid : ", u.Gettid())

	_, err := u.Getcwd(c)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(c))

	for {
	}
}
