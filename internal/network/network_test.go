package network

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/mcs-unity/ocpp-simulator/internal/event"
	"github.com/mcs-unity/ocpp-simulator/internal/record"
)

func cleanup(t *testing.T) {
	if err := os.Remove("./events.log"); err != nil {
		t.Error(err)
	}
}

func getRecord(t *testing.T) record.IRecord {
	r, err := record.New("./")
	if err != nil {
		t.Error(err)
	}
	return r
}

func TestNewNetwork(t *testing.T) {
	defer cleanup(t)
	r := getRecord(t)
	n := New(r, TIMEOUT)
	if n == nil {
		t.Error("New returned nil network")
	}
}

func TestTimeout(t *testing.T) {
	defer cleanup(t)
	r := getRecord(t)
	n := New(r, TIMEOUT)
	if err := n.SetTimeout(2 * time.Second); err == nil {
		t.Error("bypassed minimum 10 seconds validation")
	}
}

func TestGetEvent(t *testing.T) {
	defer cleanup(t)
	r := getRecord(t)
	n := New(r, TIMEOUT)
	if e := n.Event(); e == nil {
		t.Error("failed to get event handler")
	}
}

func TestNetworkTrigger(t *testing.T) {
	defer cleanup(t)
	r := getRecord(t)
	n := New(r, TIMEOUT)
	ch := make(chan event.State, 1)
	defer close(ch)
	fn := func(ctx context.Context, d event.Done, p event.Payload) {
		if v, ok := p.(event.State); ok {
			ch <- v
		}

		d <- struct{}{}
	}

	n.Event().SubScribe(event.READY, "test", fn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
