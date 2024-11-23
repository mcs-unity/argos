package action

import (
	"sync"
	"time"

	"github.com/mcs-unity/ocpp-simulator/internal/event"
	"github.com/mcs-unity/ocpp-simulator/internal/ocpp"
)

const (
	TIMEOUT = 5 * time.Second
)

type IAction interface {
	Request(ocpp.UniqueId)
	Response(ocpp.Payload, event.IEvent)
}

type IHandler interface {
	Send(a ocpp.Action) error
}

type handler struct {
	lock  sync.Locker
	event event.IEvent
	list  map[ocpp.Action]IAction
}
