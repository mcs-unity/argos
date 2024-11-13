package event

import (
	"context"
	"testing"
)

func TestNewEvent(t *testing.T) {
	c := NewEvent()
	if c == nil {
		t.Error("failed to create event handler")
	}
}

func TestSubscribeToEvent(t *testing.T) {
	c := NewEvent()

	fn := func(ctx context.Context, payload any) error {
		return nil
	}

	if err := c.SubScribe("test", fn); err != nil {
		t.Error(err)
	}
}

func TestSubscribeBadKey(t *testing.T) {
	c := NewEvent()
	if err := c.SubScribe("", nil); err == nil {
		t.Error("empty key was accepted")
	}
}

func TestSubscribeBadCallBack(t *testing.T) {
	c := NewEvent()
	if err := c.SubScribe("test", nil); err == nil {
		t.Error("nil callback was accepted")
	}
}

func TestSubscribeOrder(t *testing.T) {
	// subscribe multiple functions to event
	// trigger event according to registration order
	// get result
}

func TestDeleteSubscription(t *testing.T) {
	// remove subscription from events list
}

func TestBlockingSubscriber(t *testing.T) {
	// subscribe multiple functions to an event
	// make one function run forever
	// should stop function and continue
	// get result
}
