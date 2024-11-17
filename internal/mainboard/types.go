package mainboard

import (
	network "github.com/mcs-unity/ocpp-simulator/internal/network"
	"github.com/mcs-unity/ocpp-simulator/internal/record"
)

type IMainboard interface {
}

type mainboard struct {
	r       record.IRecord
	url     []byte
	network network.INetwork
}
