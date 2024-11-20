package connector

import (
	"os"
	"testing"

	"github.com/mcs-unity/ocpp-simulator/internal/record"
)

func cleanup(t *testing.T) {
	if err := os.Remove("./events.log"); err != nil {
		t.Error(err)
	}
}

func getRecord(t *testing.T) record.IRecord {
	r, err := record.New("./")
	if err != nil {
		t.Error(err)
	}
	return r
}

func TestNewConnector(t *testing.T) {
	defer cleanup(t)
	if _, err := New(1, getRecord(t)); err != nil {
		t.Error(err)
	}
}

func TestNewConnectorWithBadId(t *testing.T) {
	defer cleanup(t)
	if _, err := New(6, getRecord(t)); err == nil {
		t.Error("did not throw an error")
	}
}
