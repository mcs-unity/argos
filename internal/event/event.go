package event

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"sync"
	"time"
)

func (e *Event) SubScribe(k key, name string, cb Callback) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	if strings.Trim(k, "") == "" {
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

func (e *Event) Length(k key) int {
	if v, ok := e.list[k]; ok {
		return len(v)
	}

	return -1
}

func (e *Event) Remove(k key, n name) error {
	v, ok := e.list[k]
	if !ok {
		return fmt.Errorf("unable to find key %s", k)
	}

	index := -1
	for i, s := range v {
		if s.key != n {
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

func (e *Event) execute(list []subscription, payload any) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	if len(list) == 0 {
		return nil
	}

	for _, cn := range list {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		c := make(chan interface{}, 1)
		defer close(c)
		go cn.Callback(ctx, c, payload)

		select {
		case <-c:
			continue
		case <-ctx.Done():
		}
	}

	return nil
}

func (e *Event) Trigger(k key, payload any) error {
	if strings.Trim(k, "") == "" {
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
		make(map[string][]subscription),
		timeout,
	}
}
