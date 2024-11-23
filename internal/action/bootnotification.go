package action

import (
	"fmt"

	"github.com/mcs-unity/ocpp-simulator/internal/event"
	"github.com/mcs-unity/ocpp-simulator/internal/ocpp"
)

const BOOTNOTIFICATION = "BootNotification"

func (h bootNotification) Request(id ocpp.UniqueId) {
	// send message to CSMS
	fmt.Println("send message to CSMS")
}

func (h bootNotification) Response(p ocpp.Payload, e event.IEvent) {
	// requires access to the motherboard
	// here will will accept and then trigger status notification if accepted
	// if pending wait for instructions from CSMS
	// If rejected try again based on heartbeat interval
	e.Trigger(BOOTNOTIFICATION, p)
}

type bootNotification struct{}
