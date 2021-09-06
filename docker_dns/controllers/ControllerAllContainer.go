package controllers

import (
	"context"
	"log"
	"strings"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	ctx                = context.Background()
	AllContainerIPName = map[string]*DockerIPName{}
	OldContainerIPName = map[string]*DockerIPName{}
)

type AllContainer struct {
	ConfValue
	Lock sync.WaitGroup
}

func (a *AllContainer) AppendContainerValue(j types.ContainerJSON) bool {
	var ipname = &DockerIPName{}
	if j.State.Running {
		if _, ok := AllContainerIPName[j.Name]; !ok {
			ipname.ContainerIP = j.NetworkSettings.Networks[a.Donetwork].IPAddress
			ipname.ContainerName = strings.TrimPrefix(j.Name, "/")
			ipname.Status = j.State.Running
			AllContainerIPName[j.Name] = ipname
			return true
		}
	}
	if j.State.Running {
		AllContainerIPName[j.Name].ContainerIP = j.NetworkSettings.Networks[a.Donetwork].IPAddress
		AllContainerIPName[j.Name].Status = j.State.Running
		return true
	}
	return false
}

func (a *AllContainer) InspectContainer(client *client.Client, c types.Container) {
	containerjson, err := client.ContainerInspect(ctx, c.ID)
	if err != nil {
		log.Fatal(err)
	}
	if !a.AppendContainerValue(containerjson) {
		log.Printf("append %s is AllContainerIPName map error", containerjson.Name)
	}
}

func NewAllContainers() *AllContainer {
	return &AllContainer{
		ConfValue: *NewConfValue().ParseYaml(),
		Lock:      sync.WaitGroup{},
	}
}
