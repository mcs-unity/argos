package event

import (
	"context"
	"sync"
	"time"
)

const (
	TIMEOUT = 5 * time.Second
)

const (
	READY    State = "ready"
	PENDING  State = "pending"
	ACTIVE   State = "active"
	STARTING State = "starting"
	DATA     State = "data"
)

type State string
type name = string
type Payload = any
type Done = chan interface{}
type Callback func(context.Context, Done, Payload)

type IEvent interface {
	SubScribe(State, name, Callback) error
	Remove(State, name) error
	Trigger(State, Payload) error
	Length(State) int
}

type subscription struct {
	string
	Callback
}

type Event struct {
	lock    sync.Locker
	list    map[State][]subscription
	timeout time.Duration
}
