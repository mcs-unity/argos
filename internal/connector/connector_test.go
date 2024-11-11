package connector

import (
	"strconv"
	"testing"
)

func TestVerifyConnector(t *testing.T) {
	n, err := verifyConnector("1")
	if err != nil {
		t.Error(err)
	}

	if n != 1 {
		t.Errorf("expected 1 got %d", n)
	}
}

func TestVerifyBadValueConnector(t *testing.T) {
	_, err := verifyConnector("1A")
	if err == nil {
		t.Error("parsed a non digit string")
	}
}

func TestCreateConnectors(t *testing.T) {
	length := 2
	c, err := CreateConnectors(strconv.Itoa(length))
	if err != nil {
		t.Error(err)
	}

	arrLength := len(c)
	if arrLength != length {
		t.Errorf("expected array length of %d go %d", length, arrLength)
	}
}

func TestGetTransaction(t *testing.T) {
	a := GetAvailable()
	if a.GetTransaction() != -1 {
		t.Error("failed to get transaction")
	}
}

func TestStopTransaction(t *testing.T) {
	a := GetAvailable()
	transaction := 101
	if err := a.StartTransaction(transaction); err != nil {
		t.Error(err)
	}

	if err := a.StopTransaction(transaction); err != nil {
		t.Error(err)
	}

	if a.GetTransaction() != -1 {
		t.Error("failed to stop transaction")
	}
}
