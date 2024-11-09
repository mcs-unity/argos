package main

import (
	"os"

	"github.com/mcs-unity/ocpp-simulator/internal/charger"
	"github.com/mcs-unity/ocpp-simulator/internal/exception"
)

func main() {
	defer exception.Exception()
	args := os.Args

	if len(args) < 3 {
		panic("input arguments invalid please use command <websocket> <connectors>")
	}

	ch, err := charger.NewCharger([]byte(args[1]), args[2])
	if err != nil {
		panic(err)
	}

	if err := ch.Start(); err != nil {
		panic(err)
	}
}
