package controllers

type DockerIPName struct {
	ContainerName string
	ContainerIP   string
	Status        bool
}

var (
	confPath = "conf/conf.yaml"
)
