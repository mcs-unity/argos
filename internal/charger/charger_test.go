package charger

import (
	"fmt"
	"os"
	"testing"

	"github.com/mcs-unity/ocpp-simulator/internal/socket"
)

var url = []byte("ws://localhost:3000")

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
	_, err := NewCharger(url, "3", &socket.SocketMock{}, "")
	if err != nil {
		t.Error(err)
	}
}

func TestStart(t *testing.T) {
	d, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	ch, err := NewCharger(url, "3", &socket.SocketMock{}, fmt.Sprintf("%s/../../", d))
	if err != nil {
		t.Error(err)
	}

	if err := ch.Start(); err != nil {
		t.Error(err)
	}
}
