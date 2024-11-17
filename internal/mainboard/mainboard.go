package mainboard

import (
	"context"
	"fmt"

	"github.com/mcs-unity/ocpp-simulator/internal/event"
	"github.com/mcs-unity/ocpp-simulator/internal/network"
)

func (m *mainboard) networkState(ctx context.Context, d event.Done, p event.Payload) {
	v, ok := p.(event.State)
	if !ok {
		return
	}

	switch v {
	case event.READY:
		if err := m.network.Connect(m.url); err != nil {
			fmt.Println(err)
		}
	default:
		return
	}
}

func subscribe(m *mainboard, e ...event.IEvent) {
	for _, ev := range e {
		ev.SubScribe(event.READY, "connect", m.networkState)
	}
}

func New(url []byte) IMainboard {
	m := &mainboard{
		url,
		network.New(network.TIMEOUT), // load env timeout here
	}

	subscribe(m, m.network.Event())
	return m
}
