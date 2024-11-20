package connector

import (
	"errors"

	"github.com/mcs-unity/ocpp-simulator/internal/event"
)

func (s State) GetState() event.State {
	return s.state
}

func (s State) ChangePermitted(event.State) error {
	return errors.New("State is not a valid state")
}
