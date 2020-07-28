package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type Config struct {
	Credentials `ini:"credentials"`
}

type Credentials struct {
	Token string `ini:"token"`
}

func NewConfig(configName string) *Config {
	config := Config{}
	configExists := true
	if _, err := os.Stat(configName); os.IsNotExist(err) {
		configExists = false
	}
	if configExists {
		err := ini.MapTo(&config, configName)
		if err != nil {
			log.Fatalf("Fail to map config: %v", err)
		}
	}

	if config.Credentials.Token != "" {
		return &config
	}

	config.FromInput(configName)
	return &config
}

func (c *Config) FromInput(exportFileName string) {
	var token string
	fmt.Printf("Your GitHub token: ")
	fmt.Scan(&token)
	c.Credentials.Token = token
	err := c.Export(exportFileName)
	if err != nil {
		panic(err)
	}
}

func (c *Config) Export(fileName string) error {
	file := ini.Empty()
	err := file.ReflectFrom(c)
	if err != nil {
		return err
	}
	return file.SaveTo(fileName)
}
