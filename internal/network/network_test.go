package network

import (
	"context"
	"testing"
	"time"

	"github.com/mcs-unity/ocpp-simulator/internal/event"
)

func TestNewNetwork(t *testing.T) {
	n := New(TIMEOUT)
	if n == nil {
		t.Error("New returned nil network")
	}
}

func TestTimeout(t *testing.T) {
	n := New(TIMEOUT)
	if err := n.SetTimeout(2 * time.Second); err == nil {
		t.Error("bypassed minimum 10 seconds validation")
	}
}

func TestGetEvent(t *testing.T) {
	n := New(TIMEOUT)
	if e := n.Event(); e == nil {
		t.Error("failed to get event handler")
	}
}

func TestNetworkTrigger(t *testing.T) {
	n := New(TIMEOUT)
	ch := make(chan event.State, 1)
	defer close(ch)
	fn := func(ctx context.Context, d event.Done, p event.Payload) {
		if v, ok := p.(event.State); ok {
			ch <- v
		}

		d <- struct{}{}
	}

	n.Event().SubScribe(event.READY, "test", fn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		t.Error("failed to get ready state from network")
	case v := <-ch:
		if v != event.READY {
			t.Error("got bad state")
		}
	}
}
