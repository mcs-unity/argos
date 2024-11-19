package mainboard

import (
	"context"
	"fmt"

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

func subscribe(m *mainboard, s event.State, e event.IEvent, cb event.Callback, ref string) {
	m.r.Write(fmt.Sprintf("mainboard register %s events", ref), record.EVENT)
	if err := e.SubScribe(s, "connect", cb); err != nil {
		m.r.Write(err, record.ERROR)
	}
	m.r.Write(fmt.Sprintf("mainboard  %s events registered", ref), record.EVENT)
}

func New(r record.IRecord, url []byte) IMainboard {
	r.Write("mainboard booting", record.EVENT)
	m := &mainboard{
		r,
		url,
		network.New(r, network.TIMEOUT),
	}

	subscribe(m, event.READY, m.network.Event(), m.networkState, "network")
	return m
}
