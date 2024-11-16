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

func MockSubscriberWithValue(ch chan any, v int) Callback {
	return func(ctx context.Context, c Done, p Payload) {
		ch <- v
		c <- struct{}{}
	}
}

func MockPanicSubscriber() Callback {
	return func(ctx context.Context, c Done, p Payload) {
		panic("I'm a fake panic")
	}
}

func MockBlockingSubscriber() Callback {
	return func(ctx context.Context, c Done, p Payload) {
		time.Sleep(2 * time.Second)
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

func TestDeleteSubscription(t *testing.T) {
	key := "test"
	name := "testFn"
	c := NewEvent(TIMEOUT)
	fn := MockSubscriber(nil)
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
	key := "test"
	c := NewEvent(1 * time.Second)
	cn := make(chan interface{}, 2)
	defer close(cn)

	pn := MockSubscriberWithValue(cn, 0)
	if err := c.SubScribe(key, "testPanic", pn); err != nil {
		t.Error(err)
	}

	fn := MockSubscriberWithValue(cn, 1)
	if err := c.SubScribe(key, "testFn", fn); err != nil {
		t.Error(err)
	}

	if err := c.Trigger(key, key); err != nil {
		t.Error(err)
	}

	index := 0
	for i := range cn {
		if i != index {
			t.Error("invalid order")
			return
		}

		if i == 1 {
			return
		}
		index++
	}
}

func TestBlockingSubscriber(t *testing.T) {
	key := "test"
	c := NewEvent(1 * time.Second)
	cn := make(chan interface{}, 1)
	defer close(cn)

	pn := MockBlockingSubscriber()
	if err := c.SubScribe(key, "testPanic", pn); err != nil {
		t.Error(err)
	}

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

func TestSubscriberPanicking(t *testing.T) {
	key := "test"
	c := NewEvent(1 * time.Second)
	cn := make(chan interface{}, 1)
	defer close(cn)

	pn := MockPanicSubscriber()
	if err := c.SubScribe(key, "testPanic", pn); err != nil {
		t.Error(err)
	}

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
