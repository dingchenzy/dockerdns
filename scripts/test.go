package main

import (
	"fmt"
	"strings"
)

func main() {
	str := strings.TrimPrefix("/test-busybox", "/")
	fmt.Println(str)
}
