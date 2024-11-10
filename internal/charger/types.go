package charger

import (
	"sync"

	"github.com/mcs-unity/ocpp-simulator/internal/connector"
)

const (
	SOFT RebootType = "soft"
	HARD RebootType = "hard"
)

type RebootType string

type ICharger interface {
	Start() error
	Reboot(t RebootType) error
}

type Charger struct {
	lock       sync.Locker
	started    bool
	connectors []connector.IConnector
}
