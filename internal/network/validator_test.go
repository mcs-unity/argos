package network

import "testing"

func TestVerifyWebsocket(t *testing.T) {
	if err := isWebsocketProtocol([]byte("wss://localhost")); err != nil {
		t.Error(err)
	}

	if err := isWebsocketProtocol([]byte("ws://localhost")); err != nil {
		t.Error(err)
	}
}

func TestVerifyInvalidWebsocket(t *testing.T) {
	if err := isWebsocketProtocol([]byte("ws:/localhost")); err == nil {
		t.Error("failed to catch wrong protocol")
	}
}
