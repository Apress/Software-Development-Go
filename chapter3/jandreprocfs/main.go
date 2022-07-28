package main

import (
	"github.com/jandre/procfs"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strconv"
)

func main() {
	processes, _ := procfs.Processes(false)
	table := tablewriter.NewWriter(os.Stdout)

	for _, p := range processes {
		table.Append([]string{strconv.Itoa(p.Pid), p.Exe, p.Cwd})
	}
	table.Render()

	m, _ := procfs.NewMeminfo()
	log.Println(m)

}
