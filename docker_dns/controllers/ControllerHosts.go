package controllers

import (
	"fmt"
	"log"
	"os"
)

type ManagerHostFile struct {
	filepath string
	file     *os.File
}

func (m *ManagerHostFile) TrunkValue() {
	for _, v := range AllContainerIPName {
		_, err := m.file.WriteString(fmt.Sprintf("%s %s\n", v.ContainerIP, v.ContainerName))
		if err != nil {
			log.Fatal(err)
		}
	}
}

// 创建管理 hosts 文件结构体
func NewManagerManagerHostFile() *ManagerHostFile {
	fileopen, err := os.OpenFile(NewConfValue().ParseYaml().Hostfilepath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}

	return &ManagerHostFile{
		filepath: NewConfValue().ParseYaml().Hostfilepath,
		file:     fileopen,
	}
}
