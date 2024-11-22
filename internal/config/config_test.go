package config

import (
	"testing"
)

func cleanup() {
	configuration = nil
}

func TestLoadConfigFile(t *testing.T) {
	defer cleanup()
	if err := Load(); err != nil {
		t.Error(err)
	}

	if v := Get(NumberOfConnectors, UINT8); v == nil {
		t.Error("failed to load value from file")
	}
}
