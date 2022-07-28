package main

import (
	"log"
	s "syscall"
)

func main() {
	c := make([]byte, 512)

	log.Println("Getpid : ", s.Getpid())
	log.Println("Getpgrp : ", s.Getpgrp())
	log.Println("Getpgrp : ", s.Getppid())
	log.Println("Gettid : ", s.Gettid())

	_, err := s.Getcwd(c)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(c))
}
