package connector

import (
	"fmt"
	"slices"
	"strconv"
)

/*
	contains connector information
*/

func verifyConnector(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	if n > 5 {
		return 0, fmt.Errorf(
			"a charger may not have more then 5 connectors, your provided (%d) connectors",
			n,
		)
	}
	return n, nil
}

func CreateConnectors(s string) ([]IConnector, error) {
	connectors, err := verifyConnector(s)
	if err != nil {
		return nil, err
	}

	con := make([]IConnector, connectors)
	for i := 0; i < connectors; i++ {
		con[i] = GetAvailable()
	}

	return con, nil
}

func (state ConnectorState) IsNotTheSameState(s State) error {
	if state.State == s {
		return fmt.Errorf("invalid state change from %s to %s", state.State, s)
	}
	return nil
}

func (state ConnectorState) ChangePermitted(s State) bool {
	return slices.Contains(state.WhiteList, s)
}

func FetchState(s State, con ConnectorState) (IConnector, error) {
	var state IConnector
	switch s {
	case AVAILABLE:
		con.WhiteList = AvailableList
		state = &Available{con}
	case UNAVAILABLE:
		con.WhiteList = AvailableList
		state = &Unavailable{con}
	default:
		return nil, fmt.Errorf("failed to find state: %s", s)
	}

	return state, nil
}

func (state ConnectorState) ChangeState(s State, fn Callback) error {
	if err := state.IsNotTheSameState(s); err != nil {
		return err
	}

	if !state.ChangePermitted(s) {
		return fmt.Errorf("unable to transition from %s to %s", state.State, s)
	}

	newState, err := FetchState(s, state)
	if err != nil {
		return err
	}

	return fn(newState)
}

func (state *Available) StartTransaction() error {
	if state.Transaction != -1 {
		return fmt.Errorf("transaction (%d) is in progress", state.Transaction)
	}
	return nil
}

func (state *ConnectorState) Error(e ErrorCode, info string) {
	state.ErrorCode = e
	state.Info = info
}

func (state ConnectorState) Type() State {
	return state.State
}
