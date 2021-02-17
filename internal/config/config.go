package config

import "github.com/kelseyhightower/envconfig"

type Configuration struct {
	Timeout     int32  `envconfig:"timeout" required:"true" default:"15"`
	Port        int32  `envconfig:"port" required:"true" default:"8080"`
	Host        string `envconfig:"host" required:"true" default:"0.0.0.0"`
	TemplateDir string `envconfig:"template_dir" required:"true" default:"templates"`
}

var (
	EnvConfig Configuration
)

func init() {
	envconfig.Process("MPG", &EnvConfig)
}
