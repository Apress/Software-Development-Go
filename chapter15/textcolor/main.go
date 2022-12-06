package main

import "fmt"

const (
	Info  = "\x1b[1;34m%s\x1b[0m"
	Warn  = "\x1b[1;33m%s\x1b[0m"
	Error = "\x1b[1;31m%s\x1b[0m"
	Debug = "\x1b[0;36m%s\x1b[0m"
)

func main() {
	fmt.Printf(Info, "Info")
	fmt.Println("")
	fmt.Printf(Warn, "Warning")
	fmt.Println("")
	fmt.Printf(Error, "Error")
	fmt.Println("")
	fmt.Printf(Debug, "Debug")
	fmt.Println("")
}
