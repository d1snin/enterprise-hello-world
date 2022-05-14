package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Configuration struct {
	Port            string `default:"8080"`
	MessageReceiver string `default:"World"`
	Username        string `default:"admin"`
	Password        string `default:"admin"`
}

const prefix = "helloworld"

var l = log.Default()

func ParseConfiguration() *Configuration {
	var config Configuration

	if err := envconfig.Process(prefix, &config); err != nil {
		l.Fatal(err)
	}

	return &config
}
