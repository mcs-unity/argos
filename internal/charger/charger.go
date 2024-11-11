package charger

import (
	"regexp"
	"sync"
	"time"

	"github.com/mcs-unity/ocpp-simulator/internal/bootnotification"
	"github.com/mcs-unity/ocpp-simulator/internal/connector"
	"github.com/mcs-unity/ocpp-simulator/internal/socket"
)

/*
	initiate charging simulator
	contains run time information such as env
	keep configurations in file to be loaded on boot
*/

func verifyUrl(s []byte) bool {
	reg := regexp.MustCompile("^(ws://|wss://)")
	return reg.Match(s)
}

func seconds(s int) time.Duration {
	return time.Duration(s) * time.Second
}

func NewCharger(connectors string, sock socket.ISocket) (ICharger, error) {
	plugs, err := connector.CreateConnectors(connectors)
	if err != nil {
		return nil, err
	}

	return &Charger{
		&sync.Mutex{},
		false,
		plugs,
		seconds(30),
		bootnotification.REJECTED,
		sock,
	}, nil
}
