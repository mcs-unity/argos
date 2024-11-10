package charger

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (c *Charger) Start(url []byte) error {
	if !verifyUrl(url) {
		return errors.New("<websocket> argument must be either ws:// or wss:// protocol")
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	if c.started {
		return errors.New("charger is already started")
	}

	if err := c.socket.Connect(url); err != nil {
		return err
	}

	// load env config
	// begin connecting websocket
	// send BootNotification
	// retry until success

	return nil
}

func (c *Charger) die(ch chan<- interface{}) {
	if err := c.socket.Close(); err != nil {
		fmt.Println("failed to send close command")
		c.kill()
	}
	ch <- struct{}{}
}

func (c Charger) kill() {
	if err := c.socket.Terminate(); err != nil {
		fmt.Println(err)
	}
}

func (c *Charger) Stop() error {
	c.lock.Lock()
	defer c.lock.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	done := make(chan interface{}, 1)
	defer close(done)
	go c.die(done)

	select {
	case <-ctx.Done():
		return errors.New("failed to disconnect timeout error")
	case <-done:
		return nil
	}
}
