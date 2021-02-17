package config

import "testing"

func TestConfig(t *testing.T) {
	if EnvConfig.Port != 8080 {
		t.Errorf("Default config does not work")
	}
}
