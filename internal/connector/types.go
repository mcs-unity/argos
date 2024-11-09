package connector

type StatusNotification string

const (
	AVAILABLE      StatusNotification = "Available"
	PREPARING      StatusNotification = "Preparing"
	CHARGING       StatusNotification = "Charging"
	SUSPENDED_EVSE StatusNotification = "SuspendedEVSE"
	SUSPENDED_EV   StatusNotification = "SuspendedEV"
	FINISHING      StatusNotification = "Finishing"
	RESERVED       StatusNotification = "Reserved"
	UNAVAILABLE    StatusNotification = "Unavailable"
	FAULTED        StatusNotification = "Faulted"
)

type IConnector interface{}
