package record

import (
	"os"
	"testing"
)

func cleanup(t *testing.T) {
	if err := os.Remove("./events.log"); err != nil {
		t.Error(err)
	}
}

func TestNewRecord(t *testing.T) {
	defer cleanup(t)
	_, err := New("./")
	if err != nil {
		t.Error(err)
	}
}

func TestBadFilepath(t *testing.T) {
	_, err := New("")
	if err == nil {
		t.Error("failed to catch bad file path")
	}
}

func TestWriteToFile(t *testing.T) {
	defer cleanup(t)
	r, err := New("./")
	if err != nil {
		t.Error(err)
	}

	if err := r.Write("hello", EVENT); err != nil {
		t.Error(err)
	}
}
