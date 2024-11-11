package charger

import (
	"sync"
	"time"

	"github.com/mcs-unity/ocpp-simulator/internal/bootnotification"
	"github.com/mcs-unity/ocpp-simulator/internal/connector"
	"github.com/mcs-unity/ocpp-simulator/internal/socket"
)

const (
	SOFT RebootType = "soft"
	HARD RebootType = "hard"
	NONE RebootType = "none"
)

type RebootType string

type ICharger interface {
	Start() error
	Reboot(t RebootType) error
	die(ch chan<- interface{})
	Stop() error
}

// add io writer to keep track of events
// assign array of charging profiles
// add authorization cache
type Charger struct {
	url        []byte
	lock       sync.Locker
	started    bool
	connectors []connector.IConnector
	heartbeat  time.Duration // this shall be reassigned after the charger connects and performs bootNotification
	bootState  bootnotification.State
	socket     socket.ISocket
	boot       RebootType
}
