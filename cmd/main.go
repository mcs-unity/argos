package main

import (
	"os"
	"strings"

	"github.com/mcs-unity/ocpp-simulator/internal/exception"
	"github.com/mcs-unity/ocpp-simulator/internal/startup"
)

func main() {
	defer exception.Exception()
	dir, err := os.Getwd()

	if strings.Contains(dir, "cmd") {
		dir = strings.Replace(dir, "/cmd", "", 1)
	}

	if err != nil {
		panic(err)
	}
	startup.Boot(dir)
}
