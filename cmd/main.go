package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/mcs-unity/ocpp-simulator/internal/charger"
	"github.com/mcs-unity/ocpp-simulator/internal/exception"
)

func printCopyRight(path string) {
	b, err := os.ReadFile(path)
	if err != nil {
		return
	}
	fmt.Println(string(b))
	fmt.Printf("© %d MCS Unity Co.,LTD - copyright all rights reserved\n", time.Now().Year())
}

func main() {
	defer exception.Exception()
	args := os.Args

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	printCopyRight(fmt.Sprintf("%s/resources/logo.txt", dir))

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
