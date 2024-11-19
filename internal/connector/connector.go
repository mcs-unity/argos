package connector

import (
	"fmt"
)

func New(id uint8) (IConnector, error) {
	if id > 5 {
		return nil, fmt.Errorf("max value is 5 got %d", id)
	}

	return &Connector{
		id,
		&State{},
	}, nil
}
