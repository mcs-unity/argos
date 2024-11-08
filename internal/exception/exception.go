package exception

import (
	"fmt"
	"reflect"
)

/*
	add errors handling
	panic handler
	ocpp error
	other forms of errors
*/

func printError(err any) {
	errMessage := ""
	switch reflect.TypeOf(err).Name() {
	case "string":
		errMessage = err.(string)
	}

	fmt.Printf("error: %s\n", errMessage)
}

func Exception() {
	if err := recover(); err != nil {
		printError(err)
		fmt.Println("simulator has exited")
	}
}
