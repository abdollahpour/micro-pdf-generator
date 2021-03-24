package config

import "testing"

func TestConfig(t *testing.T) {
	envConfig := NewEnvConfiguration()
	if envConfig.Port != 8080 {
		t.Errorf("Default config does not work")
	}
}
