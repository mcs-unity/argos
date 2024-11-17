package network

import (
	"errors"
	"sync"
	"time"

	"github.com/mcs-unity/ocpp-simulator/internal/event"
	"github.com/mcs-unity/ocpp-simulator/internal/socket"
)

func (i *network) Connect(url []byte) error {
	if err := isWebsocketProtocol(url); err != nil {
		return err
	}

	if i.state != event.READY {
		return errors.New("unable to disconnect, state is not ready")
	}

	i.changeState(event.PENDING)
	i.lock.Lock()
	defer i.lock.Unlock()

	if err := i.sock.Connect(url); err != nil {
		i.changeState(event.READY)
		return err
	}

	i.changeState(event.ACTIVE)
	return nil
}

func (i *network) Disconnect() error {
	if i.state != event.ACTIVE {
		return errors.New("unable to disconnect, state is not active")
	}

	defer i.changeState(event.READY)
	i.changeState(event.PENDING)

	i.lock.Lock()
	defer i.lock.Unlock()

	if err := i.sock.Close(); err != nil {
		return err
	}

	return nil
}

func (i network) Event() event.IEvent {
	return i.trigger
}

func (i *network) changeState(s event.State) {
	i.state = s
	i.trigger.Trigger(s, s)
}

func New() INetwork {
	n := &network{
		&sync.Mutex{},
		event.READY,
		&socket.Socket{},
		event.NewEvent(event.TIMEOUT),
	}

	go func() {
		time.Sleep(1 * time.Second)
		n.changeState(n.state)
	}()

	return n
}
