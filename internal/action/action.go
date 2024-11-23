package action

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/mcs-unity/ocpp-simulator/internal/channel"
	"github.com/mcs-unity/ocpp-simulator/internal/event"
	"github.com/mcs-unity/ocpp-simulator/internal/ocpp"
)

var Handler IHandler

func (ac *handler) Send(a ocpp.Action) error {
	ac.lock.Lock()
	defer ac.lock.Unlock()

	id := uuid.New().String()
	c := channel.New(id)
	defer c.Close()
	handler, ok := ac.list[a]
	if !ok {
		return ocpp.OCPPError{}
	}
	handler.Request(id)
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	select {
	case <-ctx.Done():
		return fmt.Errorf("%s response exceeded timeout limit", a)
	case p := <-c.Listener():
		handler.Response(p, ac.event)
	}

	return nil
}

func init() {
	h := &handler{
		&sync.Mutex{},
		event.NewEvent(event.TIMEOUT),
		make(map[ocpp.Action]IAction, 1),
	}

	h.list[HEARTBEAT] = heartbeat{}
	h.list[BOOTNOTIFICATION] = bootNotification{}
	Handler = h
}
