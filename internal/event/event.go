package event

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/mcs-unity/ocpp-simulator/internal/exception"
)

func (e *Event) SubScribe(k State, name string, cb Callback) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	if strings.Trim(string(k), "") == "" {
		return errors.New("key may not be an empty string")
	}

	if strings.Trim(name, "") == "" {
		return errors.New("name may not be an empty string")
	}

	if cb == nil {
		return errors.New("nil pointer callback is not allowed")
	}

	e.list[k] = append(e.list[k], subscription{name, cb})
	return nil
}

func (e *Event) Length(k State) int {
	if v, ok := e.list[k]; ok {
		return len(v)
	}

	return -1
}

func (e *Event) Remove(k State, n name) error {
	v, ok := e.list[k]
	if !ok {
		return fmt.Errorf("unable to find key %s", k)
	}

	index := -1
	for i, s := range v {
		if s.string != n {
			continue
		}
		index = i
	}

	if index == -1 {
		return fmt.Errorf("unable to find callback with name %s", n)
	}

	e.list[k] = slices.Delete(v, index, 1)
	return nil
}

func runCallback(ctx context.Context, cb Callback, c Done, p Payload) {
	defer exception.Exception()
	cb(ctx, c, p)
}

func (e *Event) execute(list []subscription, p Payload) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	if len(list) == 0 {
		return nil
	}

	for _, cn := range list {
		ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
		defer cancel()

		c := make(chan interface{}, 1)
		defer close(c)
		go runCallback(ctx, cn.Callback, c, p)

		select {
		case <-c:
			continue
		case <-ctx.Done():
		}
	}

	return nil
}

func (e *Event) Trigger(k State, payload any) error {
	if strings.Trim(string(k), "") == "" {
		return errors.New("key may not be an empty string")
	}

	if v, ok := e.list[k]; ok {
		return e.execute(v, payload)
	}

	return nil
}

func NewEvent(timeout time.Duration) IEvent {
	return &Event{
		&sync.Mutex{},
		make(map[State][]subscription),
		timeout,
	}
}
