package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

// Configuration is a struct that contains the configuration.
type Configuration struct {

	// Server port.
	Port string `default:"8090"`

	// Who to refer to in the message.
	MessageReceiver string `default:"World"`

	// Username used for Basic authentication.
	Username string `default:"admin"`

	// Password used for Basic authentication.
	Password string `default:"admin"`
}

const _prefix = "helloworld"

var l = log.Default()

// ParseConfiguration parses environment variables to Configuration.
func ParseConfiguration() *Configuration {
	var config Configuration

	if err := envconfig.Process(_prefix, &config); err != nil {
		l.Fatal(err)
	}

	return &config
}
