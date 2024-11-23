package mainboard

import (
	"context"

	"github.com/mcs-unity/ocpp-simulator/internal/action"
	"github.com/mcs-unity/ocpp-simulator/internal/event"
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
	case event.ACTIVE:
		if err := m.IHandler.Send(action.BOOTNOTIFICATION); err != nil {
			m.r.Write(err, record.ERROR)
		}
	default:
		return
	}
}
