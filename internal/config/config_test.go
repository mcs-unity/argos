package config

import (
	"testing"
)

func TestLoadConfigFile(t *testing.T) {
	if err := Load("./config.json"); err != nil {
		t.Error(err)
	}

	if v := Get(NumberOfConnectors, UINT8); v == nil {
		t.Error("failed to load value from file")
	}
}

func TestLoadConfigFileTwice(t *testing.T) {
	if err := Load("./config.json"); err != nil {
		t.Error(err)
	}

	if err := Load("./config.json"); err == nil {
		t.Error("should only load once")
	}
}
