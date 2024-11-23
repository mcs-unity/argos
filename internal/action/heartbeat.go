package action

import "github.com/mcs-unity/ocpp-simulator/internal/ocpp"

const HEARTBEAT = "Heartbeat"

func (h heartbeat) Request(id ocpp.UniqueId) {
	// crate channel
	// send message with channel
	// wait for channel to trigger
}

func (h heartbeat) Response(p ocpp.Payload) {}

type heartbeat struct {
}
