package action

import (
	"context"

	"github.com/google/uuid"
	"github.com/mcs-unity/ocpp-simulator/internal/channel"
	"github.com/mcs-unity/ocpp-simulator/internal/ocpp"
)

func Submit(a ocpp.Action) error {
	id := uuid.New().String()
	c := channel.New(id)
	defer c.Close()
	handler, ok := list[a]
	if !ok {
		return ocpp.OCPPError{}
	}
	handler.Request(id)
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	select {
	case <-ctx.Done():
		//retry
		return nil
	case p := <-c.Listener():
		handler.Response(p)
	}
	return nil
}
