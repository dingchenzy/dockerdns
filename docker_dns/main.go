package main

import (
	_ "dockerdns/docker_dns/init"

	"dockerdns/docker_dns/start"
)

func main() {
	start.Go()
}
