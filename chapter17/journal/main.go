package main

import (
	j "github.com/coreos/go-systemd/v22/journal"
)

func main() {
	j.Print(j.PriErr, "PriErr - This log message is from Go application")
	j.Print(j.PriCrit, "PriCrit - This log message is from Go application")
	j.Print(j.PriAlert, "PriAlert - This log message is from Go application")
	j.Print(j.PriEmerg, "PriEmerg - This log message is from Go application")

	j.Print(j.PriWarning, "PriWarning - This log message is from Go application")
	j.Print(j.PriNotice, "PriNotice - This log message is from Go application")

	j.Print(j.PriInfo, "PriInfo - This log message is from Go application")
	j.Print(j.PriDebug, "PriDebug - This log message is from Go application")

}
