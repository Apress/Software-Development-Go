// This program should throw an OOM panic instead of being killed.

package main

import (
	_ "github.com/nogoegst/memlimit"
)

func main() {
	x := []byte{0x42}
	for {
		x = append(x, x...)
	}
}
