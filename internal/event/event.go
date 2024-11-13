package event

import (
	"context"
	"errors"
	"strings"
	"time"
)

func (e *Event) SubScribe(k key, cb Callback) error {
	if strings.Trim(k, "") == "" {
		return errors.New("key may not be an empty string")
	}

	if cb == nil {
		return errors.New("nil pointer callback is not allowed")
	}

	e.list[k] = append(e.list[k], cb)
	return nil
}

func execute(list []Callback, payload any) error {
	if len(list) == 0 {
		return nil
	}

	for _, cn := range list {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		// people might not process the context and as such we might have a dead lock
		// we should not trust users, make a wrapper
		if err := cn(ctx, payload); err != nil {
			return err
		}
	}

	return nil
}

func (e *Event) Trigger(k key, payload any) error {
	if strings.Trim(k, "") == "" {
		return errors.New("key may not be an empty string")
	}

	if v, ok := e.list[k]; ok {
		return execute(v, payload)
	}

	return nil
}

func NewEvent() IEvent {
	return &Event{
		make(map[string][]Callback),
	}
}
