package start

import (
	"context"
	"dockerdns/docker_dns/controllers"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	ctx           = context.Background()
	allcontainers = controllers.NewAllContainers()
)

func goroutine1run(host string) {
	client, err := client.NewClient("tcp://"+host, "", nil, nil)
	allcontainers := controllers.NewAllContainers()
	// 获取 docker ps 内容
	if err != nil {
		log.Fatal("conn docker err")
	}
	for {
		time.Sleep(time.Second * 1)
		containers, _ := client.ContainerList(ctx, types.ContainerListOptions{})
		// 获取容器信息，先通过 docker ps 找到容器的 id，然后基于容器 id 找到 docker inspect 内容
		for _, v := range containers {
			allcontainers.InspectContainer(client, v)
		}

		// fmt.Printf("controllersOld :::::::::::::::::: %#v\n", controllers.OldContainerIPName)

		if !controllers.Md5Run() {
			fmt.Println(controllers.Md5Run())
			controllers.OldContainerIPName = controllers.AllContainerIPName
			hostfile := controllers.NewManagerManagerHostFile()
			hostfile.TrunkValue()
		}
	}
}

func Go() {
	for _, v := range allcontainers.ConfValue.ParseYaml().Dohost {
		// fmt.Println(v)
		// allcontainers.Lock.Add(1)
		go goroutine1run(v)
	}
	allcontainers.Lock.Wait()
	for {
		time.Sleep(time.Second * 1)
		fmt.Printf("%#v\n", controllers.AllContainerIPName)
	}
}
