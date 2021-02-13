package config

import "github.com/kelseyhightower/envconfig"

type Configuration struct {
	Timeout int32 `envconfig:"timeout" required:"true" default:"15"`
	Port    int32 `envconfig:"port" required:"true" default:"8080"`
}

var (
	Config Configuration
)

func init() {
	envconfig.Process("MPG", &Config)
}
