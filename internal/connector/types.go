package connector

import (
	"github.com/mcs-unity/ocpp-simulator/internal/event"
	"github.com/mcs-unity/ocpp-simulator/internal/record"
)

const (
	AVAILABLE event.State = "AVAILABLE"
)

type IState interface {
	GetState() event.State
	ChangePermitted(event.State) error
}

type State struct {
	state event.State
}

type IConnector interface {
	ChangeState(IState)
	Event() event.IEvent
}

type Connector struct {
	connector uint8
	state     IState
	trigger   event.IEvent
	r         record.IRecord
}
