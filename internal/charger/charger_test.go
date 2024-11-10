package charger

import (
	"fmt"
	"testing"
)

func TestValidWebsocketUrl(t *testing.T) {
	if verifyUrl([]byte("wss://localhost:3000")) == false {
		t.Error("provided url is not a valid websocket url")
	}
}

func TestInvalidWebsocketUrl(t *testing.T) {
	if verifyUrl([]byte("http://localhost:3000")) == true {
		t.Error("provided url is valid but returned false ")
	}
}

func TestSeconds(t *testing.T) {
	sec := 20
	if dur := seconds(sec); dur.String() != fmt.Sprintf("%ds", sec) {
		t.Error("failed to get correct duration")
	}
}

func TestNewCharger(t *testing.T) {
	_, err := NewCharger("3")
	if err != nil {
		t.Error(err)
	}
}

func TestStart(t *testing.T) {
	ch, err := NewCharger("3")
	if err != nil {
		t.Error(err)
	}

	if err := ch.Start([]byte("ws://localhost:3000")); err != nil {
		t.Error(err)
	}
}
