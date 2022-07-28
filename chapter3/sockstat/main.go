package main

import (
	"bufio"
	"bytes"
	"github.com/lensesio/tableprinter"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	sockets = "sockets:"
	tcp     = "TCP:"
	udp     = "UDP:"
	colon   = ":"
	netstat = "/proc/net/sockstat"
)

func split(l []byte, p string) map[string]int64 {
	splitField := bytes.Fields(l)
	if len(splitField)%2 != 0 {
		return nil
	}
	m := make(map[string]int64)
	for k, v := range splitField {
		if k%2 == 1 {
			key := p + "_" + string(splitField[k-1])
			i, err := strconv.ParseInt(string(v), 10, 64)

			if err != nil {
				log.Fatalln(err)
			}
			m[key] = i
		}
	}
	return m
}

func readLine(r *bufio.Reader) ([]byte, error) {
	line, isPrefix, err := r.ReadLine()
	for isPrefix && err == nil {
		var bs []byte
		bs, isPrefix, err = r.ReadLine()
		line = append(line, bs...)
	}
	return line, err
}

func main() {
	fs, err := os.Open(netstat)
	if err != nil {
		log.Panic(err)
	}
	defer fs.Close()

	reader := bufio.NewReader(fs)
	m := make(map[string]int64)
	for {
		line, err := readLine(reader)
		if bytes.HasPrefix(line, []byte(sockets)) ||
			bytes.HasPrefix(line, []byte(tcp)) ||
			bytes.HasPrefix(line, []byte((udp))) {
			idx := bytes.Index(line, []byte((colon)))
			l := line[idx+1:]
			values := split(l, strings.ToLower(strings.TrimSpace(string(line[:idx]))))
			for k, v := range values {
				m[k] = v
			}
		}
		if err != nil {
			break
		}
	}

	printTable(m)
}

func printTable(m map[string]int64) {
	printer := tableprinter.New(os.Stdout)
	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.Print(m)
}
