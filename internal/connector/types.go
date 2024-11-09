package connector

import "time"

const (
	AVAILABLE      State = "Available"
	PREPARING      State = "Preparing"
	CHARGING       State = "Charging"
	SUSPENDED_EVSE State = "SuspendedEVSE"
	SUSPENDED_EV   State = "SuspendedEV"
	FINISHING      State = "Finishing"
	RESERVED       State = "Reserved"
	UNAVAILABLE    State = "Unavailable"
	FAULTED        State = "Faulted"
)

type State string
type ErrorCode string
type Callback = func(state IConnector) error

type IConnector interface {
	Type() State
	StartTransaction() error
	ChangeState(s State, callback Callback) error
	Error(s State, e ErrorCode, fn Callback)
}

type ConnectorState struct {
	Time        time.Time
	State       State
	Transaction int
}
