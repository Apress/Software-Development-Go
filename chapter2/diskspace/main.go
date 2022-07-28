package main

import (
	"fmt"
	syscall "golang.org/x/sys/unix"
)

const (
	gigabyte = (1024.0 * 1024.0 * 1024.0)
)

func main() {
	var statfs = syscall.Statfs_t{}
	var total uint64
	var used uint64
	var free uint64
	err := syscall.Statfs("/", &statfs)
	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err)
	} else {
		total = statfs.Blocks * uint64(statfs.Bsize)
		free = statfs.Bfree * uint64(statfs.Bsize)
		used = total - free
	}

	fmt.Printf("total Disk Space : %.1f GB\n", float64(total)/gigabyte)
	fmt.Printf("total Disk used  : %.1f GB\n", float64(used)/gigabyte)
	fmt.Printf("total Disk free  : %.1f GB\n", float64(free)/gigabyte)
}
