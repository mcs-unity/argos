package main

import (
	"os"

	"github.com/mcs-unity/ocpp-simulator/internal/exception"
)

func main() {
	defer exception.Exception()
	args := os.Args

	if len(args) < 3 {
		panic("input arguments invalid please use command <websocket> <connectors>")
	}
	/*
		1. read command line arguments
		2. validate argument length
		3. validate argument values
		4. handle errors
		5. enter loop
	*/
}
