package network

import (
	"sync"
	"time"

	"github.com/mcs-unity/ocpp-simulator/internal/event"
	"github.com/mcs-unity/ocpp-simulator/internal/record"
	"github.com/mcs-unity/ocpp-simulator/internal/socket"
)

const (
	TIMEOUT = 10 * time.Second
)

type INetwork interface {
	Connect(url []byte) error
	Disconnect() error
	SetTimeout(time.Duration) error
	Event() event.IEvent
}

type network struct {
	r       record.IRecord
	lock    sync.Locker
	state   event.State
	sock    socket.ISocket
	trigger event.IEvent
	timeout time.Duration
}
