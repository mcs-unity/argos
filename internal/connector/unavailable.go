package connector

import (
	"errors"
)

var UnAvailableList = []State{
	PREPARING,
	CHARGING,
	SUSPENDED_EV,
	SUSPENDED_EVSE,
	RESERVED,
	FAULTED,
}

func (state Unavailable) StartTransaction() error {
	return errors.New("connector is unavailable")
}

func GetUnAvailable() IConnector {
	return &Unavailable{newConnector(UNAVAILABLE, UnAvailableList)}
}
