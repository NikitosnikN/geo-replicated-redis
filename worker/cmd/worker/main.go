package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"worker/internal/config"
	"worker/internal/worker"
)

const EnvPrefix string = ""

func initConfig() (config.Config, error) {
	var c config.Config

	err := envconfig.Process(EnvPrefix, &c)

	if err != nil {
		log.Fatal(err.Error())
	}

	return c, nil
}

func main() {
	cfg, err := initConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	worker.Consume(cfg)
}
