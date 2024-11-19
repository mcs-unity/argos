package connector

import "testing"

func TestNewConnector(t *testing.T) {
	if _, err := New(1); err != nil {
		t.Error(err)
	}
}

func TestNewConnectorWithBadId(t *testing.T) {
	if _, err := New(6); err == nil {
		t.Error("did not throw an error")
	}
}
