package main

import (
	"goweb-project-seed/config"
	"goweb-project-seed/framework/core/xconfig"
	"log"
)

var configPath string

func main() {

	var cfg config.AppConfig
	err := xconfig.Load("config.yaml", &cfg)
	if err != nil {
		log.Fatalf("[start] app start failed.", err)
		return
	}

}

// func newApp(c *config.AppConfig) *boot.App
