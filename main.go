package main

import (
	_ "embed"
	"fmt"
	"os"
	"time"

	"github.com/mcs-unity/ocpp-simulator/internal/exception"
	"github.com/mcs-unity/ocpp-simulator/internal/startup"
)

//go:embed resources/logo.txt
var logo string

func printCopyRight() {
	fmt.Println(logo)
	fmt.Printf("© %d MCS Unity Co.,LTD - copyright all rights reserved\n\n", time.Now().Year())
}

func main() {
	printCopyRight()
	defer exception.Exception()
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	startup.Boot(dir)
}
