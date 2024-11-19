package network

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/mcs-unity/ocpp-simulator/internal/event"
	"github.com/mcs-unity/ocpp-simulator/internal/record"
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
	if t < 20*time.Second {
		return errors.New("timeout must be 20 seconds or greater")
	}

	n.timeout = t
	return nil
}

func (n network) Event() event.IEvent {
	return n.trigger
}

func (n *network) changeState(s event.State) {
	n.r.Write(fmt.Sprintf("network changed state from %s to %s", n.state, s), record.EVENT)
	n.state = s
	n.trigger.Trigger(s, s)
}

func New(r record.IRecord, timeout time.Duration) INetwork {
	n := &network{
		r,
		&sync.Mutex{},
		event.STARTING,
		&socket.Socket{},
		event.NewEvent(event.TIMEOUT),
		timeout,
	}

	r.Write("network booting up", record.EVENT)
	go func() {
		time.Sleep(SYSTEM_BOOT)
		n.changeState(event.READY)
	}()
	r.Write("network boot completed", record.EVENT)

	return n
}
