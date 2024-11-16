package network

import "github.com/mcs-unity/ocpp-simulator/internal/socket"

const (
	READY   state = "ready"
	PENDING state = "pending"
	ACTIVE  state = "active"
)

type state string

type INetwork interface {
	Connect(url []byte) error
	Disconnect() error
}

type network struct {
	state state
	sock  socket.ISocket
}
