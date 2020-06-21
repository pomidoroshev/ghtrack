package main

import (
	"log"

	"gopkg.in/ini.v1"
)

type Config struct {
	Credentials `ini:"credentials"`
}

type Credentials struct {
	Token string `ini:"token"`
}

func NewConfig(configName string) *Config {
	credentials := Config{}
	err := ini.MapTo(&credentials, configName)
	if err != nil {
		log.Fatalf("Fail to map config: %v", err)
	}
	return &credentials
}
