package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/mcs-unity/ocpp-simulator/internal/charger"
	"github.com/mcs-unity/ocpp-simulator/internal/exception"
)

func main() {
	defer exception.Exception()
	args := os.Args

	if len(args) < 3 {
		panic("input arguments invalid please use command <websocket> <connectors>")
	}

	ch, err := charger.NewCharger(args[2])
	if err != nil {
		panic(err)
	}

	if err := ch.Start([]byte(args[1])); err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	defer close(c)

	signal.Notify(c)

	sig := <-c

	fmt.Printf("\ngot %s exiting \n", sig.String())
	if err := ch.Stop(); err != nil {
		fmt.Println(err)
	}
}
