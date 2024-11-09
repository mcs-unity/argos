package connector

import (
	"errors"
)

func (state Unavailable) StartTransaction() error {
	return errors.New("connector is unavailable")
}
