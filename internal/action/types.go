package action

import (
	"time"

	"github.com/mcs-unity/ocpp-simulator/internal/ocpp"
)

const (
	TIMEOUT = 5 * time.Second
)

type IAction interface {
	Request(id ocpp.UniqueId)
	Response(p ocpp.Payload)
}
