package mainboard

import (
	"context"

	"github.com/mcs-unity/ocpp-simulator/internal/event"
	"github.com/mcs-unity/ocpp-simulator/internal/network"
	"github.com/mcs-unity/ocpp-simulator/internal/record"
)

func (m *mainboard) networkState(ctx context.Context, d event.Done, p event.Payload) {
	v, ok := p.(event.State)
	if !ok {
		return
	}

	switch v {
	case event.READY:
		if err := m.network.Connect(m.url); err != nil {
			m.r.Write(err, record.ERROR)
		}
	default:
		return
	}
}

func subscribe(m *mainboard, e ...event.IEvent) {
	for _, ev := range e {
		if err := ev.SubScribe(event.READY, "connect", m.networkState); err != nil {
			m.r.Write(err, record.ERROR)
		}

	}
}

func New(r record.IRecord, url []byte) IMainboard {
	m := &mainboard{
		r,
		url,
		network.New(r, network.TIMEOUT), // load env timeout here
	}

	subscribe(m, m.network.Event())
	return m
}
