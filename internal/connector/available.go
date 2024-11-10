package connector

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

func GetAvailable() IConnector {
	return &Available{newConnector(AVAILABLE, AvailableList)}
}
