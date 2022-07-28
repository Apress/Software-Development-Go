package main

import (
	"fmt"
	seccomp "github.com/seccomp/libseccomp-golang"
	"log"
	"math/rand"
	"syscall"
	"time"
)

var (
	whitelist = []string{
		"getcwd", "exit_group", "rt_sigreturn", "mkdirat", "write",
	}
	filter      *seccomp.ScmpFilter
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

	if err := configureSeccomp(); err != nil {
		log.Println(fmt.Sprintf("Failed to load seccomp filter: %v", err))
		return
	}

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

func configureSeccomp() error {
	var err error
	filter, err = seccomp.NewFilter(seccomp.ActErrno)
	if err != nil {
		return err
	}
	for _, name := range whitelist {
		syscallID, err := seccomp.GetSyscallFromName(name)
		if err != nil {
			return err
		}
		err = filter.AddRule(syscallID, seccomp.ActAllow)
		if err != nil {
			return err
		}
	}

	err = filter.Load()
	filter.Release()
	return err
}
