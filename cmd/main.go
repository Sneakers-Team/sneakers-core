package main

import (
	"flag"
	"log"
	"msqrd/sneakers/internal/apiserver"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "../configs/apiserver.toml", "path to a config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Printf("config error: %s", err)
	}

	apiserver := apiserver.NewServer(config)
	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
