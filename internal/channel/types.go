package channel

import (
	"errors"
	"sync"

	"github.com/mcs-unity/ocpp-simulator/internal/ocpp"
)

var (
	ErrClosedChannel = errors.New("channel has been closed")
	ErrNilPayload    = errors.New("payload can't be nil")
)

type IChannel interface {
	Listener() <-chan any
	Close() error
	Notify(p ocpp.Payload) error
	ID() string
}

type Channel struct {
	uuid    ocpp.UniqueId
	lock    sync.Locker
	channel chan any
	closed  bool
}
