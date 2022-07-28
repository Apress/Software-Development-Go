package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const (
	memInfo = "/proc/meminfo"
)

type memory struct {
	total uint64
	used  uint64
	free  uint64
}

type sampler struct {
	rate   time.Duration
	sample chan sample
}

type sample struct {
	memorySample memory
}

func getMemorySample() (samp memory) {
	want := map[string]bool{
		"total:": true,
		"free:":  true,
		"used:":  true}

	contents, err := ioutil.ReadFile(memInfo)
	if err != nil {
		return
	}

	reader := bufio.NewReader(bytes.NewBuffer(contents))
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fields := strings.Fields(string(line))
		fieldName := fields[0]

		_, ok := want[fieldName]
		if ok && len(fields) == 3 {
			val, numerr := strconv.ParseUint(fields[1], 10, 64)
			if numerr != nil {
				return
			}
			switch fieldName {
			case "total:":
				samp.total = val
			case "free:":
				samp.free = val
			}
		}
	}
	samp.used = samp.total - samp.free
	return
}

func (s *sampler) start() *sampler {
	s.sample = make(chan sample)
	go func() {
		for {
			var ss sample
			ss.memorySample = getMemorySample()
			s.sample <- ss
			time.Sleep(s.rate)
		}
	}()
	return s
}

func main() {
	sampler := &sampler{
		rate: 1 * time.Second,
	}
	sampler.start()

	for {
		select {
		case sampleSet := <-sampler.sample:
			s := sampleSet.memorySample
			fmt.Printf("total = %v KB, free = %v KB, used = %v KB\n",
				s.total, s.free, s.used)
		}
	}
}
