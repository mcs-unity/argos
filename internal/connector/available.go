package connector

import (
	"time"
)

var AvailableList = []State{
	PREPARING,
	CHARGING,
	SUSPENDED_EV,
	SUSPENDED_EVSE,
	RESERVED,
	UNAVAILABLE,
	FAULTED,
}

func GetAvailable() IConnector {
	return &Available{ConnectorState{time.Now(), AVAILABLE, -1, NO_ERROR, "", AvailableList}}
}
