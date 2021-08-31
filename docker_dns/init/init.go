package init

import (
	"log"
	"os"
)

func init() {
	File, err := os.OpenFile("./dockerdns.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(File)
}
