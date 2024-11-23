package channel

import (
	"context"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	c := New("123")
	defer c.Close()
	if c == nil {
		t.Error("failed to create new channel")
	}
}

func TestNewEmpty(t *testing.T) {
	c := New("")
	if c != nil {
		t.Error("should not create new channel on empty id")
	}
}

func TestChannelFlow(t *testing.T) {
	c := New("123")
	if c == nil {
		t.Error("failed to create new channel")
	}
	defer c.Close()

	value := "hello"
	go func(v string, ch IChannel) {
		ch.Notify(v)
	}(value, c)

	channel := c.Listener()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		t.Error("failed to trigger channel")
	case p := <-channel:
		if p != value {
			t.Errorf("expected %s got %s", value, p)
		}
	}
}

func TestChannelClosed(t *testing.T) {
	c := New("123")
	if err := c.Close(); err != nil {
		t.Error(err)
	}

	if err := c.Notify("test"); err == nil {
		t.Error("channel did not close")
	}
}
