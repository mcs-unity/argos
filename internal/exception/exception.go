package exception

import (
	"fmt"
)

/*
	ocpp error
	other forms of errors
*/

func printError(err any) {
	errMessage := ""

	switch v := err.(type) {
	case string:
		errMessage = v
	case error:
		errMessage = v.Error()
	default:
		errMessage = "unable to convert error to printable format"
	}

	fmt.Printf("error: %s\n", errMessage)
}

func Exception() {
	if err := recover(); err != nil {
		printError(err)
	}
}
