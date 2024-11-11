package connector

import "fmt"

var AvailableList = []State{
	AVAILABLE,
	PREPARING,
	CHARGING,
	SUSPENDED_EVSE,
	SUSPENDED_EV,
	FINISHING,
	RESERVED,
	UNAVAILABLE,
	FAULTED,
}

func (state *Available) StartTransaction(id int) error {
	if state.Transaction != -1 {
		return fmt.Errorf("transaction (%d) is in progress", state.Transaction)
	}

	state.Transaction = id
	return nil
}

func GetAvailable() IConnector {
	return &Available{newConnector(AVAILABLE, AvailableList)}
}
