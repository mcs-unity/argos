package network

import (
	"sync"

	"github.com/mcs-unity/ocpp-simulator/internal/event"
	"github.com/mcs-unity/ocpp-simulator/internal/socket"
)

type INetwork interface {
	Connect(url []byte) error
	Disconnect() error
	Event() event.IEvent
}

type network struct {
	lock    sync.Locker
	state   event.State
	sock    socket.ISocket
	trigger event.IEvent
}
