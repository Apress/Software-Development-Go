package main

import (
	"fmt"
	"github.com/google/cadvisor/client"
	v1 "github.com/google/cadvisor/info/v1"
	"github.com/sirupsen/logrus"
	"strings"
)

func Group(cinfo []v1.ContainerInfo) (map[string][]v1.ContainerInfo, error) {
	groups := make(map[string][]v1.ContainerInfo)
	for _, v := range cinfo {
		key := strings.Split(v.Name, v.Id)[0]
		_, ok := groups[key]
		if !ok {
			groups[key] = make([]v1.ContainerInfo, 1)
			groups[key][0] = v
			continue
		}
		groups[key] = append(groups[key], v)
	}
	return groups, nil
}

//run
//docker run --name my-mysql -e MYSQL_ROOT_PASSWORD=secret -v $PWD/mysql-data:/var/lib/mysql  mysql:8.0
func main() {
	cadvisor, err := client.NewClient("http://localhost:8080")

	cinfoList, err := cadvisor.AllDockerContainers(&v1.ContainerInfoRequest{})
	if err != nil {
		logrus.Errorf("read cadvisor fail, stop collect: %s", err.Error())
		return
	}

	pods, _ := Group(cinfoList)
	for key, v := range pods {
		fmt.Printf("%s: %s\n", key, v[0].Aliases[0])
	}
}
