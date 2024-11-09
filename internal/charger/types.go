package charger

import "github.com/mcs-unity/ocpp-simulator/internal/connector"

type ICharger interface {
	Start() error
}

type Charger struct {
	Connectors []connector.IConnector
}
