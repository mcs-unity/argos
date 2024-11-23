package channel

import (
	"strings"
	"sync"

	"github.com/mcs-unity/ocpp-simulator/internal/ocpp"
)

func (e *Channel) Close() error {
	e.lock.Lock()
	defer e.lock.Unlock()

	if e.closed {
		return ErrClosedChannel
	}

	close(e.channel)
	e.closed = true
	return nil
}

func (e *Channel) Notify(p ocpp.Payload) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	if e.closed {
		return ErrClosedChannel
	}

	if p == nil {
		return ErrNilPayload
	}

	e.channel <- p
	return nil
}

func (e Channel) ID() string {
	return e.uuid
}

func (e Channel) Listener() <-chan any {
	return e.channel
}

func New(id ocpp.UniqueId) IChannel {
	if strings.Trim(id, "") == "" {
		return nil
	}

	return &Channel{id, &sync.Mutex{}, make(chan any, 1), false}
}
