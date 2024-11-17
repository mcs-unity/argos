package network

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/mcs-unity/ocpp-simulator/internal/event"
	"github.com/mcs-unity/ocpp-simulator/internal/socket"
)

func (n *network) Connect(url []byte) error {
	if err := isWebsocketProtocol(url); err != nil {
		return err
	}

	if n.state != event.READY {
		return errors.New("unable to disconnect, state is not ready")
	}

	n.changeState(event.PENDING)
	n.lock.Lock()
	defer n.lock.Unlock()

	if err := n.sock.Connect(url); err != nil {
		time.Sleep(n.timeout)
		n.changeState(event.READY)
		return err
	}

	n.changeState(event.ACTIVE)
	return nil
}

func (n *network) Disconnect() error {
	if n.state != event.ACTIVE {
		return errors.New("unable to disconnect, state is not active")
	}

	defer n.changeState(event.READY)
	n.changeState(event.PENDING)

	n.lock.Lock()
	defer n.lock.Unlock()

	if err := n.sock.Close(); err != nil {
		fmt.Println(err)
		if err := n.sock.Terminate(); err != nil {
			return err
		}
	}

	return nil
}

func (n *network) SetTimeout(t time.Duration) error {
	if t < 10*time.Second {
		return errors.New("timeout must be 10 seconds or greater")
	}

	n.timeout = t
	return nil
}

func (n network) Event() event.IEvent {
	return n.trigger
}

func (n *network) changeState(s event.State) {
	n.state = s
	n.trigger.Trigger(s, s)
}

func New(timeout time.Duration) INetwork {
	n := &network{
		&sync.Mutex{},
		event.READY,
		&socket.Socket{},
		event.NewEvent(event.TIMEOUT),
		timeout,
	}

	go func() {
		time.Sleep(1 * time.Second)
		n.changeState(n.state)
	}()

	return n
}
