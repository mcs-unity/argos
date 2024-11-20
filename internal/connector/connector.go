package connector

import (
	"fmt"

	"github.com/mcs-unity/ocpp-simulator/internal/event"
	"github.com/mcs-unity/ocpp-simulator/internal/record"
)

func (e Connector) Event() event.IEvent {
	return e.trigger
}

func (e *Connector) ChangeState(s IState) {
	current, change := e.state.GetState(), s.GetState()
	e.r.Write(fmt.Sprintf("network changed state from %s to %s", current, change), record.EVENT)
	e.state = s
	e.trigger.Trigger(change, s)
}

func New(id uint8, r record.IRecord) (IConnector, error) {
	if id > 5 {
		return nil, fmt.Errorf("max value is 5 got %d", id)
	}

	return &Connector{
		id,
		&State{
			state: AVAILABLE,
		},
		event.NewEvent(event.TIMEOUT),
		r,
	}, nil
}
