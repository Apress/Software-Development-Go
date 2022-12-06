package main

import "fmt"

const (
	Underline    = "\x1b[4m"
	UnderlineOff = "\x1b[24m"
	Italics      = "\x1b[3m"
	ItalicsOff   = "\x1b[23m"
)

func main() {
	fmt.Printf("%s%s%s\n\n", Italics, "Italics text", ItalicsOff)
	fmt.Printf("%s%s%s\n\n", Underline, "Underline text", UnderlineOff)
}
