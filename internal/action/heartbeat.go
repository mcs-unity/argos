package action

import (
	"github.com/mcs-unity/ocpp-simulator/internal/event"
	"github.com/mcs-unity/ocpp-simulator/internal/ocpp"
)

const HEARTBEAT = "Heartbeat"

func (h heartbeat) Request(id ocpp.UniqueId) {
	// send message to CSMS
}

func (h heartbeat) Response(p ocpp.Payload, e event.IEvent) {
	// verify response payload before trigger
	if p == nil {
		return
	}

	e.Trigger(HEARTBEAT, p)
}

type heartbeat struct{}
