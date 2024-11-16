package network

import (
	"errors"
	"regexp"
)

func isWebsocketProtocol(url []byte) error {
	r, err := regexp.Compile("^(ws://|wss://)")
	if err != nil {
		return err
	}

	if !r.Match(url) {
		return errors.New("invalid protocol")
	}

	return nil
}
