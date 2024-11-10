package charger

import (
	"errors"
	"regexp"
	"sync"

	"github.com/mcs-unity/ocpp-simulator/internal/connector"
)

/*
	initiate charging simulator
	contains run time information such as env
	allow charger to reboot (reset connector state and follow OCPP procedure)
	keep configurations in file to be loaded on boot
*/

func verifyUrl(s []byte) bool {
	reg := regexp.MustCompile("^(ws://|wss://)")
	return reg.Match(s)
}

func (c *Charger) Start() error {
	if c.started {
		return errors.New("charger is already started")
	}
	c.started = true

	// begin connecting websocket
	return nil
}

func (c *Charger) Reboot(t RebootType) error {
	defer c.lock.Unlock()
	c.lock.Lock()

	if t == HARD {
		// simulate rebooting the hardware long reboot
		// hard reboot don't be smooth
	} else {
		// simulate rebooting the firmware short reboot
		// also known as a graceful reboot
	}

	return nil
}

func NewCharger(s []byte, n string) (ICharger, error) {
	if !verifyUrl(s) {
		return nil, errors.New("<websocket> argument must be either ws:// or wss:// protocol")
	}

	connectors, err := connector.CreateConnectors(n)
	if err != nil {
		return nil, err
	}

	return &Charger{&sync.Mutex{}, false, connectors}, nil
}
