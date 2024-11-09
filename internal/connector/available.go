package connector

import (
	"fmt"
)

func (state *Available) ChangeState(s State, fn Callback) error {
	if err := state.IsNotTheSameState(s); err != nil {
		return err
	}

	if s == FINISHING {
		return fmt.Errorf("unable to transition from %s to %s", state.State, s)
	}

	newState, err := FetchState(s, state.ConnectorState)
	if err != nil {
		return err
	}

	return fn(newState)
}

type Available struct {
	ConnectorState
}
