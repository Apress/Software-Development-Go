package main

import (
	"fmt"
	"github.com/google/cadvisor/client"
	v1 "github.com/google/cadvisor/info/v1"
	"github.com/sirupsen/logrus"
)

func main() {
	cadvisor, err := client.NewClient("http://localhost:8080")

	cinfoList, err := cadvisor.ContainerInfo("/", &v1.ContainerInfoRequest{})
	if err != nil {
		logrus.Errorf("read cadvisor fail, stop collect: %s", err.Error())
		return
	}

	for _, c := range cinfoList.Subcontainers {
		fmt.Printf("%s \n", c.Name)
	}
}
