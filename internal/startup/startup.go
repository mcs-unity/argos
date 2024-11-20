package startup

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mcs-unity/ocpp-simulator/internal/mainboard"
	"github.com/mcs-unity/ocpp-simulator/internal/record"
)

func Boot(dir string) {
	args := os.Args

	if len(args) < 2 {
		panic("input arguments invalid please use command <websocket> <connectors>")
	}

	r, err := record.New(dir)
	if err != nil {
		panic(err)
	}

	mainboard.New(r, []byte(args[1]))

	c := make(chan os.Signal, 1)
	defer close(c)

	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT)
	sig := <-c

	fmt.Printf("\ngot %s exiting \n", sig.String())
}
