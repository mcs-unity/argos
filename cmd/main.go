package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mcs-unity/ocpp-simulator/internal/exception"
	"github.com/mcs-unity/ocpp-simulator/internal/mainboard"
	"github.com/mcs-unity/ocpp-simulator/internal/record"
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
