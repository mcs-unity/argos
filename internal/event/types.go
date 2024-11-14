package event

import (
	"context"
	"sync"
	"time"
)

const (
	TIMEOUT = 30 * time.Second
)

type key = string
type name = string
type Payload = any
type Done = chan interface{}
type Callback func(context.Context, Done, Payload)

type IEvent interface {
	SubScribe(key, name, Callback) error
	Remove(key, name) error
	Trigger(key, Payload) error
	Length(key) int
}

type subscription struct {
	key
	Callback
}

type Event struct {
	lock    sync.Locker
	list    map[key][]subscription
	timeout time.Duration
}
