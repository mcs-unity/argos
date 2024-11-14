package event

import (
	"context"
	"testing"
	"time"
)

func MockSubscriber(ch chan any) Callback {
	return func(ctx context.Context, c Done, p Payload) {
		ch <- p
		c <- struct{}{}
	}
}

func MockSimpleSubscriber() Callback {
	return func(ctx context.Context, c Done, p Payload) {
		c <- struct{}{}
	}
}

func TestNewEvent(t *testing.T) {
	c := NewEvent(TIMEOUT)
	if c == nil {
		t.Error("failed to create event handler")
	}
}

func TestSubscribeToEvent(t *testing.T) {
	c := NewEvent(TIMEOUT)

	if err := c.SubScribe("test", "testFn", MockSubscriber(make(chan any))); err != nil {
		t.Error(err)
	}
}

func TestSubscribeBadKey(t *testing.T) {
	c := NewEvent(TIMEOUT)
	if err := c.SubScribe("", "testFn", nil); err == nil {
		t.Error("empty key was accepted")
	}
}

func TestSubscribeBadName(t *testing.T) {
	c := NewEvent(TIMEOUT)
	if err := c.SubScribe("test", "", nil); err == nil {
		t.Error("empty name was accepted")
	}
}

func TestSubscribeBadCallBack(t *testing.T) {
	c := NewEvent(TIMEOUT)
	if err := c.SubScribe("test", "testFn", nil); err == nil {
		t.Error("nil callback was accepted")
	}
}

func TestSubscriber(t *testing.T) {
	key := "test"
	c := NewEvent(TIMEOUT)
	cn := make(chan interface{}, 1)
	defer close(cn)
	fn := MockSubscriber(cn)
	if err := c.SubScribe(key, "testFn", fn); err != nil {
		t.Error(err)
	}

	if err := c.Trigger(key, key); err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case v := <-cn:
		if v != key {
			t.Errorf("expected %s got %s", key, v)
		}
	case <-ctx.Done():
		t.Error("failed to trigger within timeout")
	}
}

func TestSubscribeOrder(t *testing.T) {
	// subscribe multiple functions to event
	// trigger event according to registration order
	// get result
}

func TestDeleteSubscription(t *testing.T) {
	key := "test"
	name := "testFn"
	c := NewEvent(TIMEOUT)
	fn := MockSimpleSubscriber()
	if err := c.SubScribe(key, name, fn); err != nil {
		t.Error(err)
	}

	if err := c.Remove(key, name); err != nil {
		t.Error(err)
	}

	count := c.Length(key)
	if count > 0 {
		t.Errorf("exacted -1 got %d", count)
	}
}

func TestSubscriberTriggeredBasedOnOrder(t *testing.T) {
	// subscribe multiple functions to an event
	// make sure subs are executed based on subscription order
	// get result
}

func TestBlockingSubscriber(t *testing.T) {
	// subscribe multiple functions to an event
	// make one function run forever
	// should stop function and continue
	// get result
}

func TestSubscriberPanicking(t *testing.T) {
	// panic within a function
	// should stop function and continue
	// get result
}
