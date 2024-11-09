package connector

import (
	"errors"
	"fmt"
)

func (state Unavailable) StartTransaction() error {
	return errors.New("connector is unavailable")
}

func (state *Unavailable) ChangeState(s State, fn Callback) error {
	if err := state.IsNotTheSameState(s); err != nil {
		return err
	}

	if s != AVAILABLE {
		return fmt.Errorf("unable to transition from %s to %s", state.State, s)
	}

	newState, err := FetchState(s, state.ConnectorState)
	if err != nil {
		return err
	}

	return fn(newState)
}

type Unavailable struct {
	ConnectorState
}
