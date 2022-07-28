package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	appArmorEnabledPath = "/sys/module/apparmor/parameters/enabled"
	appArmorModePath    = "/sys/module/apparmor/parameters/mode"
)

func appArmorMode() (mode string) {
	content, err := ioutil.ReadFile(appArmorModePath)
	if err != nil {
		log.Fatalln(err)
		return
	}
	return strings.TrimSpace(string(content))
}

func appArmorEnabled() (support bool) {
	content, err := ioutil.ReadFile(appArmorEnabledPath)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") { // not found means not support
			return
		}
		log.Fatalln(err)
		return
	}
	return strings.TrimSpace(string(content)) == "Y"
}

func main() {
	fmt.Println("AppArmor mode : ", appArmorMode())
	fmt.Println("AppArmor is enabled : ", appArmorEnabled())
}
