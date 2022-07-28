package main

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	tcpEstablished = 1
	netTcpFile     = "/proc/net/tcp"
)

// establishedTcpConns returns a map containing the number of established tcp connections.
func establishedTcpConns() (map[string]int, error) {
	inos, err := readFdDir(os.Getpid())
	if err != nil {
		return nil, err
	}

	tcpConns, err := netTcpConns()
	if err != nil {
		return nil, err
	}

	conns := make(map[string]int)
	for rHost, inodes := range tcpConns {
		for _, ino := range inodes {
			if _, ok := inos[ino]; ok {
				conns[rHost]++
			}
		}
	}
	return conns, nil
}

// netTcpConns reads the tcp established connections and inodes from /proc/net/tcp
func netTcpConns() (map[string][]int, error) {
	f, err := os.Open(netTcpFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	r.ReadBytes('\n') // strip header

	cache := make(map[string][]int)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		parts := strings.Fields(strings.TrimSpace(line))
		if len(parts) < 11 {
			continue
		}

		state, err := strconv.ParseInt(parts[3], 16, 32)
		if err != nil {
			continue
		}

		if state != tcpEstablished {

		}

		remoteHost := strings.Split(parts[2], ":")
		if len(remoteHost) < 2 {
			continue
		}

		var remotePort int
		_, err = fmt.Sscanf(remoteHost[1], "%04X", &remotePort)
		if err != nil {
			continue
		}

		remoteAddr, err := hex.DecodeString(remoteHost[0])
		if err != nil {
			continue
		}

		ino, err := strconv.Atoi(parts[9])
		if err != nil {
			continue
		}

		// Correct the ordering
		binary.BigEndian.PutUint32(remoteAddr, binary.LittleEndian.Uint32(remoteAddr))

		rHost := fmt.Sprintf("%v:%d", net.IP(remoteAddr), remotePort)
		cache[rHost] = append(cache[rHost], ino)
	}

	return cache, nil
}

// readFdDir reads the file descriptors for a process and returns a inode to fd mapping.
// The map can then be used to backtrace open connections.
func readFdDir(pid int) (map[int]int, error) {
	fdDir := fmt.Sprintf("/proc/%d/fd", pid)
	f, err := os.Open(fdDir)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fis, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	socks := make(map[int]int)
	for _, fi := range fis {
		i, err := strconv.Atoi(fi.Name())
		if err != nil {
			continue
		}

		lnk, err := os.Readlink(fmt.Sprintf("%s/%s", fdDir, fi.Name()))
		if err != nil {
			continue
		}

		if !strings.HasPrefix(lnk, "socket:[") {
			continue
		}
		s := strings.TrimPrefix(lnk, "socket:[")
		ino, err := strconv.Atoi(s[:len(s)-1])
		if err != nil {
			continue
		}
		socks[ino] = i
	}

	return socks, nil
}
