package main

import "os"
import "flag"
import "fmt"
import "strings"
import "debug/elf"
import "syscall"
//import "debug/dwarf"

func dump_dynstr(file *elf.File) {
	fmt.Printf("DynStrings:\n")
	dynstrs, _ := file.DynString(elf.DT_NEEDED)
	for _, e := range dynstrs {
		fmt.Printf("\t%s\n", e)
	}
	dynstrs, _ = file.DynString(elf.DT_SONAME)
	for _, e := range dynstrs {
		fmt.Printf("\t%s\n", e)
	}
	dynstrs, _ = file.DynString(elf.DT_RPATH)
	for _, e := range dynstrs {
		fmt.Printf("\t%s\n", e)
	}
	dynstrs, _ = file.DynString(elf.DT_RUNPATH)
	for _, e := range dynstrs {
		fmt.Printf("\t%s\n", e)
	}
}

func dump_symbols(file *elf.File) {
	fmt.Printf("Symbols:\n")
	symbols, _ := file.Symbols()
	for _, e := range symbols {
		if !strings.EqualFold(e.Name, "") {
			fmt.Printf("\t%s\n", e.Name)
		}
	}
}

func dump_elf(filename string) int {
	file, err := elf.Open(filename)
	if err != nil {
		fmt.Printf("Couldn’t open file : \"%s\" as an ELF.\n")
		return 2
	}
	dump_dynstr(file)
	dump_symbols(file)
	return 0
}

func init_debug(filename string) int {
	attr := &os.ProcAttr{ Sys: &syscall.SysProcAttr{ Ptrace: true } }
	if proc, err := os.StartProcess(filename, []string { "/" }, attr); err == nil {
		proc.Wait()
		foo := syscall.PtraceAttach(proc.Pid)
		fmt.Printf("Started New Process: %v.\n", proc.Pid)
		fmt.Printf("PtraceAttach res: %v.\n", foo)
		return 0
	}
	return 2;
}

func main() {
	if len(os.Args) > 1 {
		filename := flag.String("filename", "", "A binary ELF file.")
		action := flag.String("action", "", "Action to make: {dump|debug}.")
		flag.Parse()
		if *filename == "" || *action == "" {
			goto Usage
		}
		file, err := os.Stat(*filename)
		if os.IsNotExist(err) {
			fmt.Printf("No such file or directory: %s.\n", *filename)
			goto Error
		} else if mode := file.Mode(); mode.IsDir() {
			fmt.Printf("Parameter must be a file, not a " +
			"directory.\n")
			goto Error
		}
		f, err := os.Open(*filename)
		if err != nil {
			fmt.Printf("Couldn’t open file: \"%s\".\n", *filename)
			goto Error
		}
		f.Close()
		fmt.Printf("Tracing program : \"%s\".\n", *filename)
		fmt.Printf("Action : \"%s\".\n", *action)
		switch *action {
		case "debug": os.Exit(init_debug(*filename))
		case "dump": os.Exit(dump_elf(*filename))
		}
	} else {
		goto Usage
	}

	Usage:
		fmt.Printf("Usage of ./main:\n" +
			"  -action=\"{dump|debug}\": Action to make.\n" +
			"  -filename=\"file\": A binary ELF file.\n")
		goto Error

	Error:
		os.Exit(2)
}
