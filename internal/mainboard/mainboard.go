package mainboard

import (
	"fmt"

	"github.com/mcs-unity/ocpp-simulator/internal/action"
	"github.com/mcs-unity/ocpp-simulator/internal/event"
	"github.com/mcs-unity/ocpp-simulator/internal/network"
	"github.com/mcs-unity/ocpp-simulator/internal/record"
)

func subscribe(m *mainboard, e event.IEvent, cb event.Callback, ref string, s ...event.State) {
	m.r.Write(fmt.Sprintf("mainboard register %s events", ref), record.EVENT)
	for _, st := range s {
		if err := e.SubScribe(st, "connect", cb); err != nil {
			m.r.Write(err, record.ERROR)
		}
	}

	m.r.Write(fmt.Sprintf("mainboard  %s events registered", ref), record.EVENT)
}

func New(r record.IRecord, url []byte) IMainboard {
	r.Write("mainboard booting", record.EVENT)
	m := &mainboard{
		r,
		url,
		network.New(r, network.TIMEOUT),
		action.Handler,
	}

	subscribe(m, m.network.Event(), m.networkState, "network", event.READY, event.ACTIVE)
	return m
}
