package connector

import (
	"fmt"
	"strconv"
)

/*
	contains connector information
	implement state management
*/

func verifyConnector(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	if n > 5 {
		return 0, fmt.Errorf(
			"a charger may not have more then 5 connectors, your provided (%d) connectors",
			n,
		)
	}
	return n, nil
}

func CreateConnectors(s string) ([]IConnector, error) {
	connectors, err := verifyConnector(s)
	if err != nil {
		return nil, err
	}

	con := make([]IConnector, connectors)
	for i := 0; i < connectors; i++ {
		con[i] = Unavailable{}
	}

	return con, nil
}
