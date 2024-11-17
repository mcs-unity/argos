package record

import "io"

const (
	ERROR       = "error"
	EVENT       = "event"
	TRANSACTION = "transaction"
)

type event string
type payload any

type IRecord interface {
	Write(payload, event) error
}

type Record struct {
	io.Writer
}
