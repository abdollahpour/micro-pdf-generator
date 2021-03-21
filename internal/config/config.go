package config

import (
	"io/ioutil"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	Timeout int    `envconfig:"timeout" required:"true" default:"15"`
	Port    int32  `envconfig:"port" required:"true" default:"8080"`
	Host    string `envconfig:"host" required:"true" default:"0.0.0.0"`
	TempDir string `envconfig:"temp_dir" required:"true"`
	MaxSize int    `envconfig:"max_size" required:"true" default:"6"`
}

var (
	EnvConfig Configuration
)

func init() {
	envconfig.Process("MPG", &EnvConfig)
	if len(EnvConfig.TempDir) == 0 {
		dir, err := ioutil.TempDir("", "*.html")
		if err != nil {
			log.Fatal(err)
		}
		EnvConfig.TempDir = dir
	}
}
