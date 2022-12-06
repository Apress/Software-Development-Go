package main

import (
	"bytes"
	"log"
	ex "os/exec"
)

func main() {
	log.Println("--------------")
	log.Println("Running ip link")
	log.Println("--------------")
	Run("ip link")

	log.Println("---------------")
	log.Println("Running noexist")
	log.Println("---------------")
	Run("noexist")

	log.Println("Running uname -r")
	log.Println("----------------")
	Run("uname -r")
}

func Run(arg string) {
	var cmd *ex.Cmd
	cmd = ex.Command("/bin/sh", "-c", arg)
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	if err := cmd.Run(); err != nil {
		log.Println("%v", err)
		return
	}

	log.Println(stdoutBuf.String())
	log.Println(stderrBuf.String())
}
