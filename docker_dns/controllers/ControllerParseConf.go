package controllers

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type ConfValue struct {
	Dohost       []string `yaml:"dohost"`
	Donetwork    string   `yaml:"donetwork"`
	Hostfilepath string   `yaml:"hostfilepath"`
}

func (c *ConfValue) ParseYaml() *ConfValue {
	filevalue, err := os.ReadFile(confPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(filevalue, c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func NewConfValue() *ConfValue {
	return &ConfValue{}
}
