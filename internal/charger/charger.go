package charger

import (
	"errors"
	"regexp"

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

func (s *Charger) Start() error {
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

	return &Charger{connectors}, nil
}
