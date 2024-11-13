package event

import "context"

type key = string
type Callback func(ctx context.Context, payload any) error

type IEvent interface {
	SubScribe(k key, cb Callback) error
	Trigger(k key, payload any) error
}

type Event struct {
	list map[key][]Callback
}
