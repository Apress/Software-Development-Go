// memlimit_linux.go - set memory rlimit to cgroup memory limit
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to this module of memlimit, using the creative
// commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

// +build linux

package memlimit

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"syscall"
)

const memoryLimitCgroupLocation = "/sys/fs/cgroup/memory/memory.limit_in_bytes"

func readFileString(filename string) (string, error) {
	d, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(d), "\n"), nil
}

func init() {
	limitStr, err := readFileString(memoryLimitCgroupLocation)
	if err != nil {
		// not in a container?
		fmt.Fprintf(os.Stderr, "memlimit not applied: %v", err)
		return
	}
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		panic(err)
	}
	rlimit := syscall.Rlimit{
		Cur: limit,
		Max: limit,
	}
	err = syscall.Setrlimit(syscall.RLIMIT_AS, &rlimit)
	if err != nil {
		panic(err)
	}
}
