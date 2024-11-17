package mainboard

import network "github.com/mcs-unity/ocpp-simulator/internal/network"

type IMainboard interface {
}

type mainboard struct {
	url     []byte
	network network.INetwork
}
